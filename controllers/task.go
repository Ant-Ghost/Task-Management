package controllers

import (
	"net/http"
	"strconv"
	"zocket/task_manager/dtos"
	"zocket/task_manager/services"
	"zocket/task_manager/utils"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var input dtos.CreateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid input", err,
		)
		return
	}

	currentUserId, exists := c.Get("userId")
	if !exists {
		utils.ErrorResponseFormatter(
			c, http.StatusUnauthorized, "User not authenticated", nil,
		)
		return
	}

	currentUserIdUint, ok := currentUserId.(uint)
	if !ok {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Invalid user ID", nil,
		)
		return
	}

	task, err := services.CreateTask(input, &currentUserIdUint)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to create task", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusCreated, "Task created successfully", task,
	)
}
func GetTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid task ID", err,
		)
		return
	}

	task, err := services.GetTaskByID(uint(id))
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusNotFound, "Task not found", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "Task fetched successfully", task,
	)
}
func GetAllTasks(c *gin.Context) {
	tasks, err := services.GetAllTasks()
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to fetch tasks", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "Tasks fetched successfully", tasks,
	)
}
func UpdateTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid task ID", err,
		)
		return
	}

	var input dtos.UpdateTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid input", err,
		)
		return
	}

	task, err := services.UpdateTask(uint(id), input)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to update task", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "Task updated successfully", task,
	)
}
func DeleteTask(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid task ID", err,
		)
		return
	}

	if err := services.DeleteTask(uint(id)); err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to delete task", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "Task deleted successfully", nil,
	)
}
func GetTasksByFilters(c *gin.Context) {
	var input dtos.FilterTaskRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid input", err,
		)
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err == nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err == nil {
		pageSize = 10
	}

	tasks, err := services.GetTasksByFilters(input, page, pageSize)
	if err != nil {
		// utils.TraceError(err)
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to fetch tasks", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "Tasks fetched successfully", tasks,
	)
}
func GetAllSubTasks(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid task ID", err,
		)
		return
	}

	tasks, err := services.GetSubTasksByParentId(uint(id))
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to fetch subtasks", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "Subtasks fetched successfully", tasks,
	)
}

func GetAiSuggestion(c *gin.Context) {
	var input dtos.GetAiSuggestion
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid input", err,
		)
		return
	}

	suggestion, err := services.GetTaskBreakUpByAi(input)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to get AI suggestion", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "AI suggestion fetched successfully", suggestion,
	)
}
