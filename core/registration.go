package core

import (
	"errors"

	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/secure/core/roletype"
)

type Registration struct {
	App            Application
	Name           string
	Email          string
	Password       string
	PasswordRepeat string
}

func Register(r Registration) (husk.Recorder, error) {
	if r.Password != r.PasswordRepeat {
		return nil, errors.New("passwords do not match")
	}

	if len(r.App.Name) == 0 {
		return nil, errors.New("application name can not be empty")
	}

	if len(r.App.InstanceID) == 0 {
		return nil, errors.New("instance id can not be empty")
	}

	if emailExists(r.Email) {
		return nil, errors.New("email already in use")
	}

	user, err := NewUser(r.Name, r.Email)

	if err != nil {
		return nil, err
	}

	user.SecurePassword(r.Password)
	user.AddTrace(getRegistrationTrace(r))
	user.AddRole(r.App.Name, roletype.User)

	rec := ctx.Users.Create(user)
	defer ctx.Users.Save()

	if rec.Error != nil {
		return nil, rec.Error
	}

	return rec.Record, nil
}
