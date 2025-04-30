package slack

import (
	"fmt"
	"zocket/task_manager/models"
	"zocket/task_manager/repository"
)

type SlackEventType int

const (
	TaskAssigned SlackEventType = iota
	TaskUnAssigned
	TaskDetailsUpdated
	TaskPriorityUpdated
	TaskStatusUpdated
	TaskDeleted
)

func (event SlackEventType) String() string {
	switch event {
	case TaskAssigned:
		return "Task Assigned"
	case TaskUnAssigned:
		return "Task Unassigned"
	case TaskDetailsUpdated:
		return "Task Details Updated"
	case TaskPriorityUpdated:
		return "Task Priority Updated"
	case TaskStatusUpdated:
		return "Task Status Updated"
	case TaskDeleted:
		return "Task Deleted"
	default:
		return "Unknown Event"
	}
}

type SlackEvent struct {
	Task        *models.Task
	UserSlackID string
	EventType   SlackEventType
}

func formatMessage(task *models.Task, action string) string {

	titleString := fmt.Sprintf("Title: %s", task.Title)
	descriptionString := fmt.Sprintf("Description: %s", task.Description)
	statusString := fmt.Sprintf("Status: %s", task.Status)
	priorityString := fmt.Sprintf("Priority: %s", task.Priority)
	eventString := fmt.Sprintf("Action: %s", action)

	message := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s",
		titleString,
		descriptionString,
		statusString,
		priorityString,
		eventString,
	)

	return message
}

func (s SlackEvent) SendSlackMessage() {
	message := formatMessage(s.Task, s.EventType.String())
	SendDirectMessage(s.UserSlackID, message)
}

func SlackTasksHandler(eventType SlackEventType, tasks [](models.Task)) {

	for _, task := range tasks {
		user, err := repository.GetUserByID(*task.AssignToID)
		if err == nil {
			slackEvent := SlackEvent{
				Task:        &task,
				UserSlackID: user.SlackID,
				EventType:   eventType,
			}
			slackEvent.SendSlackMessage()
		}
	}
}
