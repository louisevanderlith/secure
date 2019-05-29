package core

import (
	comms "github.com/louisevanderlith/comms/core"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/mango"
)

func SendRegistrationEmail(usr User, token string, instanceID string) error {
	msg := comms.Message{
		Body:         "https://auth....",
		Name:         usr.Name,
		To:           usr.Email,
		Email:        "System",
		TemplateName: "registration.html",
	}

	var contain interface{}
	_, err := mango.DoSEND("POST", token, &contain, instanceID, "Comms.API", "message", msg)

	return err
}

func SendResetRequestEmail(usr User, forgotKey husk.Key, token string, instanceID string) error {
	msg := comms.Message{
		Body:         fmt.Sprintf("https://"),
		Name:         usr.Name,
		To:           usr.Email,
		Email:        "System",
		TemplateName: "resetrequest.html",
	}

	var contain interface{}
	_, err := mango.DoSEND("POST", token, &contain, instanceID, "Comms.API", "message", msg)

	return err
}
