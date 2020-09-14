package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/kong/prime"
	"strings"
)

type profileCalc func(result interface{}, obj prime.Profile) error

func (f profileCalc) Map(result interface{}, obj hsk.Record) error {
	return f(result, obj.GetValue().(prime.Profile))
}

func Whitelist(prefix string) profileCalc {
	return func(result interface{}, obj prime.Profile) error {
		lst := result.(*[]string)

		for _, clnt := range obj.Clients {
			for _, rsrc := range clnt.AllowedResources {
				if strings.HasPrefix(rsrc, prefix) {
					*lst = append(*lst, clnt.Url)
				}
			}
		}

		return nil
	}
}
