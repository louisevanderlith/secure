package core

import (
	"github.com/louisevanderlith/kong/prime"
	"log"
	"time"

	"github.com/louisevanderlith/husk"
)

type SafeUser struct {
	Key         husk.Key
	Name        string
	Verified    bool
	DateCreated time.Time
}

func createSafeUser(user husk.Recorder) SafeUser {
	data := user.Data().(prime.User)
	meta := user.Meta()

	result := SafeUser{
		Key:         meta.Key,
		Name:        data.Name,
		Verified:    data.Verified,
		DateCreated: time.Unix(0, meta.Key.Stamp),
	}

	return result
}

func GetUsers(page, size int) []SafeUser {
	var result []SafeUser
	users, err := ctx.Users.Find(page, size, husk.Everything())

	if err != nil {
		log.Println(err)
		return nil
	}

	itor := users.GetEnumerator()

	for itor.MoveNext() {
		currUser := itor.Current()

		sfeUser := createSafeUser(currUser)
		result = append(result, sfeUser)
	}

	return result
}
