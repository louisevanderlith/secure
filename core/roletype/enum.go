package roletype

import (
	"strings"
)

type Enum = int

const (
	Admin Enum = iota
	Owner
	User
	Unknown
)

var roleTypes = [...]string{
	"Admin",
	"Owner",
	"User",
	"Unknown"}

func GetRoleType(name string) Enum {
	var result Enum

	for k, v := range roleTypes {
		if strings.ToUpper(name) == v {
			result = Enum(k)
			break
		}
	}

	return result
}
