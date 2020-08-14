package core

import (
	"errors"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"strings"
)

type context struct {
	Profiles  husk.Tabler
	Resources husk.Tabler
}

var ctx context

func Context() context {
	return ctx
}

func (c context) GetProfile(id string) (prime.Profile, error) {
	rec, err := c.Profiles.FindFirst(byID(id))

	if err != nil {
		return prime.Profile{}, err
	}

	return rec.Data().(prime.Profile), nil
}

func (c context) GetResource(name string) (prime.Resource, error) {
	rec, err := c.Resources.FindFirst(byName(name))

	if err != nil {
		return prime.Resource{}, err
	}

	return rec.Data().(prime.Resource), nil
}

//GetWhitelist will return a list of registered domains which may call this service
func (c context) GetWhitelist(prefix string) []string {
	var lst []string
	err := c.Profiles.Calculate(&lst, Whitelist(prefix))

	if err != nil {
		log.Println("GetWhitelist", err)
		return nil
	}

	return lst
}

func (c context) GetProfileClient(id string) (prime.Profile, prime.Client, error) {
	idparts := strings.Split(id, ".")

	if len(idparts) != 2 {
		return prime.Profile{}, prime.Client{}, errors.New("id is invalid")
	}

	prof, err := c.GetProfile(idparts[0])

	if err != nil {
		return prime.Profile{}, prime.Client{}, err
	}

	clnt, err := prof.GetClient(idparts[1])

	if err != nil {
		return prime.Profile{}, prime.Client{}, err
	}

	return prof, clnt, nil
}

func (c context) UpdateProfile(k husk.Key, p prime.Profile) error {
	obj, err := ctx.Profiles.FindByKey(k)

	if err != nil {
		return err
	}

	err = obj.Set(p)

	if err != nil {
		return err
	}

	err = ctx.Profiles.Update(obj)

	if err != nil {
		return err
	}

	return ctx.Profiles.Save()
}

func (c context) UpdateResource(k husk.Key, p prime.Resource) error {
	obj, err := ctx.Resources.FindByKey(k)

	if err != nil {
		return err
	}

	err = obj.Set(p)

	if err != nil {
		return err
	}

	err = ctx.Resources.Update(obj)

	if err != nil {
		return err
	}

	return ctx.Resources.Save()
}

func CreateContext() {
	defer seed()
	ctx = context{
		Profiles:  husk.NewTable(prime.Profile{}),
		Resources: husk.NewTable(prime.Resource{}),
	}
}

func Shutdown() {
	ctx.Profiles.Save()
	ctx.Resources.Save()
}

func seed() {
	err := ctx.Profiles.Seed("db/profiles.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Profiles.Save()

	err = ctx.Resources.Seed("db/resources.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Resources.Save()
}
