package core

import "github.com/louisevanderlith/secure/core/roletype"

//ActionMap maps URL Actions [GET, POST, PUT, DELETE] to required RoleType
type ActionMap map[string]roletype.Enum
