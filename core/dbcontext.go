package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Users     husk.Tabler
	Clients   husk.Tabler
	Forgotten husk.Tabler
}

var ctx context

func CreateContext() {
	defer seed()

	ctx = context{
		Users:     husk.NewTable(User{}),
		Clients:   husk.NewTable(Client{}),
		Forgotten: husk.NewTable(Forgot{}),
	}
}

func Shutdown() {
	ctx.Users.Save()
	ctx.Clients.Save()
	ctx.Forgotten.Save()
}

func seed() {
	err := ctx.Users.Seed("db/users.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Users.Save()

	err = ctx.Clients.Seed("db/clients.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Clients.Save()
}
