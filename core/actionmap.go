package core

import "github.com/louisevanderlith/droxolite/roletype"

//ActionMap maps URL Actions [GET, POST, PUT, DELETE] to required RoleType
type ActionMap map[string]roletype.Enum
