package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Users     husk.Tabler
	Forgotten husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Users:     husk.NewTable(new(User)),
		Forgotten: husk.NewTable(new(Forgot)),
	}
}

func Shutdown() {
	ctx.Users.Save()
	ctx.Forgotten.Save()
}

func seed() {
	err := ctx.Users.Seed("db/users.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Users.Save()
}
