package core

import (
	"errors"
	"fmt"

	"github.com/louisevanderlith/husk"
)

//Forgot is used to keep a list of Users which have requested to Change their passwords.
type Forgot struct {
	UserKey  husk.Key
	Redeemed bool
}

func (v Forgot) Valid() error {
	return husk.ValidateStruct(&v)
}

//ResetRequest When users forget their passwords, we create a redeemable 'Reset Request' which can be used to reset their password.
//returns the Request Link or an error
func RequestReset(email, host string) (string, error) {
	rec, err := ctx.Users.FindFirst(emailFilter(email))

	if err != nil {
		return "", err
	}

	forget := Forgot{
		UserKey:  rec.GetKey(),
		Redeemed: false,
	}

	forgt, err := ctx.Forgotten.Create(forget)

	if err != nil {
		return "", err
	}

	resetLink := fmt.Sprintf("%s/%s", host, forgt.GetKey())

	return resetLink, nil
}

func ResetPassword(forgotKey husk.Key, password string) error {
	rec, err := ctx.Forgotten.FindByKey(forgotKey)

	if err != nil {
		return err
	}

	forgetData := rec.Data().(Forgot)

	if forgetData.Redeemed {
		return errors.New("already redeemed")
	}

	if len(password) < 6 {
		return errors.New("password length must be 6 or more characters")
	}

	_, err = ctx.Users.FindByKey(forgetData.UserKey)

	if err != nil {
		return err
	}

	//Change the Users password
	//usrRec.SecurePassword(password)

	//Redeem the Forgot
	forgetData.Redeemed = true

	err = ctx.Users.Save()

	if err != nil {
		return err
	}

	err = ctx.Forgotten.Save()

	if err != nil {
		return err
	}

	return nil
}
