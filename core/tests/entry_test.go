package secure_test

import (
	"github.com/louisevanderlith/kong/prime"
	"testing"

	uuid "github.com/nu7hatch/gouuid"

	"github.com/louisevanderlith/entity/core"
)

func getFakeApp() core.Application {
	instID, _ := uuid.NewV4()
	result := core.Application{
		InstanceID: instID.String(),
		IP:         "127.0.0.1",
		Location:   "-26.1496832, 28.035481599999997",
		Name:       "Entry.TEST",
	}

	return result
}

func TestRegistration_Good_Pass(t *testing.T) {
	r := core.Registration{
		App:            getFakeApp(),
		Email:          "joe@fake.com",
		Name:           "Joe",
		Password:       "w34k###",
		PasswordRepeat: "w34k###",
	}

	rec, err := core.Register(r)

	if err != nil {
		t.Error(err)
	}

	data := rec.Data().(core.User)

	if len(data.Roles) == 0 {
		t.Error("No Roles")
	}
}

func TestRegistration_Bad_Fail(t *testing.T) {
	r := core.Registration{
		App:      getFakeApp(),
		Name:     "Joe",
		Password: "w34k###",
	}

	rec, err := core.Register(r)

	if err == nil {
		t.Log("no error found error")
		t.Error(rec)
	}
}

func TestRegistration_Ugly_Fail(t *testing.T) {
	r := core.Registration{
		App:            getFakeApp(),
		Email:          "joe%fake.com",
		Name:           "Joe",
		Password:       "w34k!c$651d",
		PasswordRepeat: "w34k!c$651d",
	}

	rec, err := core.Register(r)

	if err == nil {
		t.Log("no error found error")
		t.Error(rec)
	}
}

func TestLogin_Good_Pass(t *testing.T) {
	r := core.Registration{
		App:            getFakeApp(),
		Email:          "joe@fake.com",
		Name:           "Joe",
		Password:       "w34k###",
		PasswordRepeat: "w34k###",
	}

	_, err := core.Register(r)

	if err != nil {
		t.Error(err)
		return
	}

	if err != nil && err.Error() != "email already in use" {
		t.Error(err)
		return
	}

	authreq := core.Authentication{
		App:      getFakeApp(),
		Email:    "joe@fake.com",
		Password: "w34k###",
	}

	_, err = core.Login(authreq)

	if err != nil {
		t.Error(err)
	}
}

func TestLogin_Bad_Fail(t *testing.T) {

}

func TestLogin_Ugly_Fail(t *testing.T) {

}

func TestSecurePassword_EasywaytoGenerateHash(t *testing.T) {
	input := "J2tXZoPMhhUyV4Y3"
	usr := prime.NewUser("Timmy", "timmy@fake.com", input, )

	t.Log(usr.Password)
	t.Fail()
}
