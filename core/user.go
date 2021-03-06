package core

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/secure/core/tracetype"

	"github.com/louisevanderlith/husk"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name        string `hsk:"size(75)"`
	Verified    bool   `hsk:"default(false)"`
	Email       string `hsk:"size(128)"`
	Password    string `hsk:"min(6)"`
	LoginDate   time.Time
	LoginTraces []LoginTrace
	Roles       []Role
}

func (u User) Valid() (bool, error) {
	valid, common := husk.ValidateStruct(&u)

	if !valid {
		return false, common
	}

	if !strings.Contains(u.Email, "@") {
		return false, errors.New("email is invalid")
	}

	return true, nil
}

func NewUser(name, email string) (*User, error) {
	result := new(User)
	result.Name = name
	result.Email = email
	result.Verified = false

	return result, nil
}

func GetUser(key husk.Key) (*User, error) {
	rec, err := ctx.Users.FindByKey(key)

	if err != nil {
		return nil, err
	}

	return rec.Data().(*User), nil
}

func (u *User) SecurePassword(plainPassword string) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(plainPassword), cost)

	if err != nil {
		log.Print("securePassword: ", err)
	}

	u.Password = string(hashedPwd)
}

func UpdateRoles(key husk.Key, roles []Role) error {
	obj, err := ctx.Users.FindByKey(key)

	if err != nil {
		return err
	}

	c := obj.Data().(*User)
	c.Roles = roles

	err = obj.Set(c)

	if err != nil {
		return err
	}

	defer ctx.Users.Save()
	return ctx.Users.Update(obj)
}

func (u *User) AddRole(appName string, role roletype.Enum) {
	appRole := Role{appName, role}
	u.Roles = append(u.Roles, appRole)
}

func (u *User) AddTrace(trace LoginTrace) {
	if trace.TraceEnv == tracetype.Login {
		u.LoginDate = time.Now()
	}

	u.LoginTraces = append(u.LoginTraces, trace)
}

func (u *User) RoleMap() map[string]roletype.Enum {
	result := make(map[string]roletype.Enum)

	for _, v := range u.Roles {
		result[v.ApplicationName] = v.Description
	}

	return result
}

func getUsers(page, size int) husk.Collection {
	return ctx.Users.Find(page, size, husk.Everything())
}

func getUser(email string) (husk.Recorder, error) {
	return ctx.Users.FindFirst(emailFilter(email))
}

func emailExists(email string) bool {
	return ctx.Users.Exists(emailFilter(email))
}
