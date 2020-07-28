package core

import (
	"github.com/louisevanderlith/husk"
)

type Role struct {
	ApplicationName string
	Description     int
}

func (o Role) Valid() error {
	return husk.ValidateStruct(o)
}
