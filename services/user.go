package services

import (
	"zocket/task_manager/dtos"
	"zocket/task_manager/models"
	"zocket/task_manager/repository"
	"zocket/task_manager/utils"
)

func RegisterUser(input dtos.RegisterUserRequest) (*models.User, error) {

	// Generate a hashed password
	hashedPassword, err := utils.GenerateHash(input.Password)
	if err != nil {
		return nil, err
	}

	// Create a new user instance
	user := &models.User{
		Username:   input.Username,
		Email:      input.Email,
		Password:   hashedPassword,
		SlackID:    input.SlackID,
		Experience: input.Experience,
	}

	return repository.CreateUser(user)
}

func GetUserByID(id uint) (*models.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	user.Password = "" // Don't return the password
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(id uint, input dtos.UpdateUserRequest) (*models.User, error) {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	// Update the user fields
	if input.Username != "" {
		user.Username = input.Username
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.SlackID != "" {
		user.SlackID = input.SlackID
	}
	if input.Experience != nil {
		user.Experience = *(input.Experience)
	}

	return repository.UpdateUser(user)
}

func DeleteUser(id uint) error {
	err := repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(input dtos.LoginUserRequest) (*string, error) {
	user, err := repository.GetUserByUsername(input.Username)
	if err != nil {
		return nil, err
	}

	// Check if the password is correct
	if !utils.MatchPassword(user.Password, input.Password) {
		return nil, utils.ErrInvalidCredentials
	}

	// Create a jwt token for the user
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
