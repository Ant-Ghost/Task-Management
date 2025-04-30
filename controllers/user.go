package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"zocket/task_manager/dtos"
	"zocket/task_manager/services"
	"zocket/task_manager/utils"
)

func Register(c *gin.Context) {
	var input dtos.RegisterUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid input", err,
		)
		return
	}

	_, err := services.RegisterUser(input)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to register user", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusCreated, "User registered successfully", nil,
	)
}
func Login(c *gin.Context) {
	var input dtos.LoginUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid input", err,
		)
		return
	}

	token, err := services.LoginUser(input)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusUnauthorized, "Invalid credentials", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "Login successful", gin.H{"token": token},
	)

}
func GetUser(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid user ID", err,
		)
		return
	}

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusNotFound, "User not found", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "User fetched successfully", user,
	)
}
func GetMyUser(c *gin.Context) {

	value, exists := c.Get("userId")
	if !exists {
		utils.ErrorResponseFormatter(
			c, http.StatusUnauthorized, "User ID not found in context", nil,
		)
		return
	}

	id, ok := value.(uint)
	if !ok {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Invalid user ID type", nil,
		)
		return
	}

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusNotFound, "User not found", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "User fetched successfully", user,
	)
}
func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to fetch users", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "Users fetched successfully", users,
	)
}
func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid user ID", err,
		)
		return
	}

	var input dtos.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid input", err,
		)
		return
	}

	user, err := services.UpdateUser(uint(id), input)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to update user", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "User updated successfully", user,
	)
}
func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusBadRequest, "Invalid user ID", err,
		)
		return
	}

	if err := services.DeleteUser(uint(id)); err != nil {
		utils.ErrorResponseFormatter(
			c, http.StatusInternalServerError, "Failed to delete user", err,
		)
		return
	}

	utils.SuccessResponseFormatter(
		c, http.StatusOK, "User deleted successfully", nil,
	)
}
