package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/secure/core/roletype"
)

type Role struct {
	ApplicationName string
	Description     roletype.Enum
}

func (o Role) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
