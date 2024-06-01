package service

import (
	"backend/dao/model"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(username, email, password string) error {
	user := &model.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := s.DB.Create(user).Error
	return err
}

// GetUserByID retrieves a user by their ID
func (s *UserService) GetUserByID(id int) (*model.User, error) {
	var user model.User
	result := s.DB.First(&user, id)
	return &user, result.Error
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(user *model.User) error {
	return s.DB.Save(user).Error
}

// DeleteUser deletes a user by their ID
func (s *UserService) DeleteUser(id int) error {
	return s.DB.Delete(&model.User{}, id).Error
}

// GetUserByUsernameAndPassword 根据用户名和密码获取用户ID
func (s *UserService) GetUserByUsernameAndPassword(username, password string) (int, error) {
	var user model.User
	result := s.DB.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.UserID, nil
}

// UserExists checks if a user exists by username or email
func (s *UserService) UserExists(username, email string) (bool, error) {
	var user model.User

	// Check for a user by username or email
	result := s.DB.Where("username = ? OR email = ?", username, email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// No record found, user does not exist
			return false, nil
		}
		// Another error occurred
		return false, result.Error
	}

	// User found
	return true, nil
}
