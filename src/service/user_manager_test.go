package service_test

import (
	"testing"

	"github.com/manujas/meli_go_course/src/domain"
	"github.com/manujas/meli_go_course/src/service"
)

func TestRegisterAUserAndRememberIt(t *testing.T) {

	// Inicialization
	var user *domain.User
	var err error

	name := "Emanuel"
	nick := "manujas"
	pass := "123456"
	mail := "manujas@mail.com"

	// Operation
	user = domain.NewUser(name, nick, mail, pass)
	err = service.RegisterNewUser(user)

	// Validation
	if err != nil {
		t.Error("No error expected")
	}

	if length := len(service.RegisteredUsers); length != 1 {
		t.Errorf("Expected one(1) user registered, but gets %v", length)
	}

	service.RegisteredUsers = service.RegisteredUsers[0:0]
}

func TestRegisterNewUserAndTheDataItsOk(t *testing.T) {
	// Inicialization
	var user *domain.User
	var err error

	name := "Emanuel"
	nick := "manujas"
	pass := "123456"
	mail := "manujas@mail.com"

	// Operation
	user = domain.NewUser(name, nick, mail, pass)
	err = service.RegisterNewUser(user)

	// Validation
	if err != nil {
		t.Error("No error expected")
	}

	if n := service.RegisteredUsers[0].Name; n != name {
		t.Errorf("Expected New User Name of %v and get %v", name, n)
	}

	if ni := service.RegisteredUsers[0].Nick; ni != nick {
		t.Errorf("Expected New User Nick of %v and get %v", nick, ni)
	}

	if m := service.RegisteredUsers[0].Mail; m != mail {
		t.Errorf("Expected New User Mail of %v and get %v", mail, m)
	}

	if p := service.RegisteredUsers[0].Pass; p != pass {
		t.Errorf("Expected New User Pass of %v and get %v", pass, p)
	}

	service.RegisteredUsers = service.RegisteredUsers[0:0]
}

func TestRegisterAUserWithoutNameGetsError(t *testing.T) {

	// Inicialization
	var user *domain.User
	var err error

	name := ""
	nick := "manujas"
	pass := "123456"
	mail := "manujas@mail.com"

	// Operation
	user = domain.NewUser(name, nick, mail, pass)
	err = service.RegisterNewUser(user)

	// Validation
	if err != nil && err.Error() != "Name is required" {
		t.Errorf("Expected 'Name is required' error message, but gets %v", err.Error())
	}

	if length := len(service.RegisteredUsers); length != 0 {
		t.Errorf("Expected zero(0) user registered, but get %v", length)
	}
}

func TestRegisterAUserWithoutNIckGetsError(t *testing.T) {

	// Inicialization
	var user *domain.User
	var err error

	name := "Emanuel"
	nick := ""
	pass := "123456"
	mail := "manujas@mail.com"

	// Operation
	user = domain.NewUser(name, nick, mail, pass)
	err = service.RegisterNewUser(user)

	// Validation
	if err != nil && err.Error() != "Nick is required" {
		t.Errorf("Expected 'Nick is required' error message, but gets %v", err.Error())
	}

	if length := len(service.RegisteredUsers); length != 0 {
		t.Errorf("Expected zero(0) user registered, but get %v", length)
	}
}

func TestRegisterAUserWithoutMailGetsError(t *testing.T) {

	// Inicialization
	var user *domain.User
	var err error

	name := "Emanuel"
	nick := "manujas"
	pass := "123456"
	mail := ""

	// Operation
	user = domain.NewUser(name, nick, mail, pass)
	err = service.RegisterNewUser(user)

	// Validation
	if err != nil && err.Error() != "Mail is required" {
		t.Errorf("Expected 'Mail is required' error message, but gets %v", err.Error())
	}

	if length := len(service.RegisteredUsers); length != 0 {
		t.Errorf("Expected zero(0) user registered, but get %v", length)
	}
}

func TestRegisterAUserWithoutPassGetsError(t *testing.T) {

	// Inicialization
	var user *domain.User
	var err error

	name := "Emanuel"
	nick := "manujas"
	pass := ""
	mail := "manujas@mail.com"

	// Operation
	user = domain.NewUser(name, nick, mail, pass)
	err = service.RegisterNewUser(user)

	// Validation
	if err != nil && err.Error() != "Pass is required" {
		t.Errorf("Expected 'Pass is required' error message, but gets %v", err.Error())
	}

	if length := len(service.RegisteredUsers); length != 0 {
		t.Errorf("Expected zero(0) user registered, but get %v", length)
	}
}
