package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/prime"
)

type resourceFilter func(obj prime.Resource) bool

func (f resourceFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(prime.Resource))
}

//byName filter will filter by client_id
func byName(name string) resourceFilter {
	return func(obj prime.Resource) bool {
		return obj.Name == name
	}
}
