package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/prime"
)

type userFilter func(obj prime.User) bool

func (f userFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(prime.User))
}

//Email filter will filter by email and verification status
func emailFilter(email string) userFilter {
	return func(obj prime.User) bool {
		return obj.Email == email && obj.Verified
	}
}
