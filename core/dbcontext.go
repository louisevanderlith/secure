package core

import (
	"encoding/json"
	"errors"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"os"
	"reflect"
	"strings"
)

type context struct {
	Profiles  husk.Table
	Resources husk.Table
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
	err := c.Profiles.Map(&lst, Whitelist(prefix))

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

func (c context) UpdateProfile(k hsk.Key, p prime.Profile) error {
	return ctx.Profiles.Update(k, p)
}

func (c context) UpdateResource(k hsk.Key, p prime.Resource) error {
	return ctx.Resources.Update(k, p)
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
	profiles, err := profileSeed()

	if err != nil {
		panic(err)
	}

	err = ctx.Profiles.Seed(profiles)

	if err != nil {
		panic(err)
	}

	resources, err := resourceSeed()

	if err != nil {
		panic(err)
	}

	err = ctx.Resources.Seed(resources)

	if err != nil {
		panic(err)
	}
}

func profileSeed() (collections.Enumerable, error) {
	f, err := os.Open("db/profiles.seed.json")

	if err != nil {
		return nil, err
	}

	var items []prime.Profile
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}

func resourceSeed() (collections.Enumerable, error) {
	f, err := os.Open("db/resources.seed.json")

	if err != nil {
		return nil, err
	}

	var items []prime.Resource
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}
