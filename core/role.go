package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango/enums"
)

type Role struct {
	ApplicationName string
	Description     enums.RoleType
}

func (o Role) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}
