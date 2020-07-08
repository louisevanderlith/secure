package core

import (
	"errors"
	"github.com/louisevanderlith/kong/prime"

	"github.com/louisevanderlith/husk"
)

type Registration struct {
	Name           string
	Email          string
	Password       string
	PasswordRepeat string
	ProfileClient  string
}

func Register(r Registration) (husk.Recorder, error) {
	if r.Password != r.PasswordRepeat {
		return nil, errors.New("passwords do not match")
	}

	if ctx.Users.Exists(emailFilter(r.Email)) {
		return nil, errors.New("email already in use")
	}

	contc := prime.Contacts{
		{
			Icon:  "fa-mail",
			Name:  "email",
			Value: r.Email,
		},
	}

	ctx.GetProfile(r.ProfileClient)

	//TODO: Make dynamic
	//Should provide only basic Resources, the rest will be unlocked later
	user := prime.NewUser(r.Name, r.Email, r.Password, false, contc, nil)

	rec, err := ctx.Users.Create(user.(prime.User))

	if err != nil {
		return nil, err
	}

	err = ctx.Users.Save()

	if err != nil {
		return nil, err
	}

	return rec, nil
}
