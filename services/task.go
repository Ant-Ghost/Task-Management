package services

import (
	"fmt"
	"zocket/task_manager/dtos"
	"zocket/task_manager/models"
	"zocket/task_manager/openai"
	"zocket/task_manager/repository"
	"zocket/task_manager/slack"
	"zocket/task_manager/websocket"
)

func CreateTask(input dtos.CreateTaskRequest, currentUserId *uint) (*models.Task, error) {
	// Create a new task instance

	task := &models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      models.Pending.String(),
		AssignToID:  input.AssignToID,
		CreatedByID: currentUserId,
	}

	if input.ParentID != nil {
		task.ParentID = input.ParentID
	}

	if input.Priority == "" {
		task.Priority = models.Low.String()
	}

	if input.ReportToID == nil {
		task.ReportToID = currentUserId
	}

	newTask, err := repository.CreateTask(task)

	if err == nil {

		go websocket.EmitEvent(
			websocket.TaskCreated,
			task.ID,
			task.ParentID,
		)

		slack.SlackTasksHandler(
			slack.TaskAssigned,
			[](models.Task){*task},
		)

	}

	return newTask, err
}

func CreateManyTasks(input dtos.CreateManyTasksRequest, currentUserId *uint) ([]models.Task, error) {
	tasks := make([]models.Task, len(input.Tasks))

	for i, taskInput := range input.Tasks {
		tasks[i] = models.Task{
			Title:       taskInput.Title,
			Description: taskInput.Description,
			Status:      models.Pending.String(),
			AssignToID:  taskInput.AssignToID,
			CreatedByID: currentUserId,
			Priority:    taskInput.Priority,
		}
		if taskInput.ParentID != nil {
			tasks[i].ParentID = taskInput.ParentID
		}
	}

	newTasks, err := repository.CreateManyTasks(tasks)
	if err == nil {
		go websocket.EmitEvent(
			websocket.TaskCreated,
			tasks[0].ID,
			tasks[0].ParentID,
		)

		go slack.SlackTasksHandler(
			slack.TaskAssigned,
			tasks,
		)
	}

	return newTasks, err
}

func GetTaskByID(id uint) (*models.Task, error) {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}
func GetAllTasks() ([]models.Task, error) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func UpdateTask(id uint, input dtos.UpdateTaskRequest) (*models.Task, error) {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	// Update the task fields
	if input.Title != "" {
		task.Title = input.Title
	}
	if input.Description != "" {
		task.Description = input.Description
	}

	updatedTask, err := repository.UpdateTask(task)
	if err == nil {
		go websocket.EmitEvent(
			websocket.TaskUpdated,
			task.ID,
			task.ParentID,
		)

		if input.Title != "" || input.Description != "" {
			go slack.SlackTasksHandler(
				slack.TaskDetailsUpdated,
				[]models.Task{*task},
			)
		}

	}

	return updatedTask, err
}

func UpdatePriority(id uint, input dtos.UpdateTaskPriorityRequest) (*models.Task, error) {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	// Update the task priority
	task.Priority = input.Priority

	updatedTask, err := repository.UpdateTask(task)
	if err == nil {
		go websocket.EmitEvent(
			websocket.TaskUpdated,
			task.ID,
			task.ParentID,
		)

		go slack.SlackTasksHandler(
			slack.TaskPriorityUpdated,
			[]models.Task{*task},
		)
	}

	return updatedTask, err
}

func UpdateStatus(id uint, input dtos.UpdateTaskStatusRequest) (*models.Task, error) {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	// Update the task status
	task.Status = input.Status

	updatedTask, err := repository.UpdateTask(task)
	if err == nil {
		go websocket.EmitEvent(
			websocket.TaskUpdated,
			task.ID,
			task.ParentID,
		)

		go slack.SlackTasksHandler(
			slack.TaskStatusUpdated,
			[]models.Task{*task},
		)
	}

	return updatedTask, err
}

func UpdateAssignTo(id uint, input dtos.UpdateTaskAssignToRequest) (*models.Task, error) {
	task, err := repository.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	oldTask := *task

	newAssignedUser, err := repository.GetUserByID(*input.AssignToID)
	if err != nil {
		return nil, err
	}

	task.AssignToID = &newAssignedUser.ID

	updatedTask, err := repository.UpdateTask(task)
	if err == nil {
		go websocket.EmitEvent(
			websocket.TaskUpdated,
			task.ID,
			task.ParentID,
		)
		go slack.SlackTasksHandler(
			slack.TaskAssigned,
			[]models.Task{*updatedTask},
		)

		go slack.SlackTasksHandler(
			slack.TaskUnAssigned,
			[]models.Task{oldTask},
		)
	}

	return updatedTask, err
}

func DeleteTask(id uint) error {
	existingTask, err := repository.GetTaskByID(id)
	if err != nil {
		return err
	}

	if err := repository.DeleteTask(existingTask.ID); err != nil {
		return err
	}

	go websocket.EmitEvent(
		websocket.TaskDeleted,
		existingTask.ID,
		existingTask.ParentID,
	)

	go slack.SlackTasksHandler(
		slack.TaskDeleted,
		[]models.Task{*existingTask},
	)

	return nil
}
func GetTasksByFilters(input dtos.FilterTaskRequest, page, pageSize int) ([]models.Task, error) {

	filters := make(map[string]interface{})
	if input.Status != "" {
		filters["status"] = input.Status
	}
	if input.ParentID != nil {
		filters["parent_id"] = input.ParentID
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	tasks, err := repository.GetTasksByFilters(
		filters,
		input.SearchQuery,
		page, pageSize,
	)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetSubTasksByParentId(parentID uint) ([]models.Task, error) {
	tasks, err := repository.GetTasksByParentID(parentID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetTaskBreakUpByAi(input dtos.GetAiSuggestion) (string, error) {
	// Call the OpenAI API to get the task breakup
	titleString := fmt.Sprintf("Title: %s", input.Title)
	descriptionString := fmt.Sprintf("Description: %s", input.Description)
	contextString := fmt.Sprintf("Context: %s", input.Context)

	prompt := fmt.Sprintf(
		"%s\n%s\n%s",
		titleString,
		descriptionString,
		contextString,
	)

	response, err := openai.GetAIResponse(prompt)
	if err != nil {
		return "", err
	}
	return response, nil
}
