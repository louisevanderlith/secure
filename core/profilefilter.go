package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/kong/prime"
	"strings"
)

type profileFilter func(obj prime.Profile) bool

func (f profileFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(prime.Profile))
}

//byID filter will filter by client_id
func byID(id string) profileFilter {
	return func(obj prime.Profile) bool {
		return strings.ToLower(obj.Title) == id
	}
}
