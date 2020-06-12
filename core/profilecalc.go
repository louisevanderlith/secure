package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/prime"
)

type profileCalc func(result interface{}, obj prime.Profile) error

func (f profileCalc) Calc(result interface{}, obj husk.Dataer) error {
	return f(result, obj.(prime.Profile))
}

func Whitelist() profileCalc {
	return func(result interface{}, obj prime.Profile) error {
		lst := result.(*[]string)

		for _, clnt := range obj.Clients {
			*lst = append(*lst, clnt.Url)
		}

		return nil
	}
}
