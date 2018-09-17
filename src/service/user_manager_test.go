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

	if len(service.RegisteredUsers) != 1 {
		t.Error("Expected one(1) user registered")
	}
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
}
