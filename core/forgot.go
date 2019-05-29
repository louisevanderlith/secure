package core

import (
	"errors"

	"github.com/louisevanderlith/husk"
)

//Forgot is used to keep a list of Users which have requested to Change their passwords.
type Forgot struct {
	UserKey  husk.Key
	Redeemed bool
}

func (v Forgot) Valid() (bool, error) {
	return husk.ValidateStruct(&v)
}

func RequestReset(email, token, instanceID string) error {
	rec, err := getUser(email)

	if err != nil {
		return err
	}

	forget := Forgot{
		UserKey:  rec.GetKey(),
		Redeemed: false,
	}

	cset := ctx.Forgotten.Create(forget)

	if cset.Error != nil {
		return cset.Error
	}

	return SendResetRequestEmail(*rec.Data().(*User), cset.Record.GetKey(), token, instanceID)
}

func ResetPassword(forgotKey husk.Key, password string) error {
	rec, err := ctx.Forgotten.FindByKey(forgotKey)

	if err != nil {
		return err
	}

	forgetData := rec.Data().(*Forgot)

	if forgetData.Redeemed {
		return errors.New("already redeemed")
	}

	if len(password) < 6 {
		return errors.New("password length must be 6 or more characters")
	}

	usrRec, err := GetUser(forgetData.UserKey)

	if err != nil {
		return err
	}

	//Change the Users password
	usrRec.SecurePassword(password)

	//Redeem the Forgot
	forgetData.Redeemed = true

	ctx.Users.Save()
	ctx.Forgotten.Save()

	return nil
}
