package service

import (
	"fmt"

	"github.com/manujas/meli_go_course/src/domain"
)

// RegisteredUsers slice
var RegisteredUsers = make([]domain.User, 0)

// RegisterNewUser create a new user
func RegisterNewUser(user *domain.User) error {

	if user == nil {
		return fmt.Errorf("Expected a User")
	}
	if user.Nick == "" {
		return fmt.Errorf("Nick is required")
	}

	if user.Name == "" {
		return fmt.Errorf("Name is required")
	}

	if user.Mail == "" {
		return fmt.Errorf("Mail is required")
	}

	if user.Pass == "" {
		return fmt.Errorf("Pass is required")
	}

	RegisteredUsers = append(RegisteredUsers, *user)

	return nil
}
