package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Users husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Users: husk.NewTable(new(User)),
	}
}

func Shutdown() {
	ctx.Users.Save()
}

func seed() {
	err := ctx.Users.Seed("db/users.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Users.Save()
}
