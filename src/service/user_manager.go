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
	RegisteredUsers = append(RegisteredUsers, *user)

	return nil
}
