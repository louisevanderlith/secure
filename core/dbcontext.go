package core

import (
	"errors"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
	"github.com/louisevanderlith/kong/prime"
	"log"
	"strings"
)

type context struct {
	Users     husk.Tabler
	Profiles  husk.Tabler
	Resources husk.Tabler
	Forgotten husk.Tabler
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
func (c context) GetWhitelist() []string {
	var lst []string
	err := c.Profiles.Calculate(&lst, Whitelist())

	if err != nil {
		log.Println("GetWhitelist", err)
		return nil
	}

	return lst
}

func (c context) GetUser(id string) prime.Userer {
	k, err := husk.ParseKey(id)

	if err != nil {
		return nil
	}

	rec, err := c.Users.FindByKey(k)

	if err != nil {
		return nil
	}

	return rec.Data().(prime.User)
}

func (c context) GetUserByName(username string) (string, prime.Userer) {
	rec, err := c.Users.FindFirst(emailFilter(username))

	if err != nil {
		return "", nil
	}

	return rec.GetKey().String(), rec.Data().(prime.User)
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

func CreateContext() {
	defer seed()
	gobbr := serials.GobSerial{}
	ctx = context{
		Users:     husk.NewTable(prime.User{}, gobbr),
		Profiles:  husk.NewTable(prime.Profile{}, gobbr),
		Resources: husk.NewTable(prime.Resource{}, gobbr),
		Forgotten: husk.NewTable(Forgot{}, gobbr),
	}
}

func Shutdown() {
	ctx.Users.Save()
	ctx.Profiles.Save()
	ctx.Resources.Save()
	ctx.Forgotten.Save()
}

func seed() {
	/*err := ctx.Users.Seed("db/users.seed.json")

	if err != nil {
		panic(err)
	}

	ctx.Users.Save()*/

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
