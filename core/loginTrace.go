package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/secure/core/tracetype"
)

type LoginTrace struct {
	Location        string `hsk:"null;size(128)"`
	IP              string `hsk:"null;size(50)"`
	Allowed         bool   `hsk:"default(true)"`
	InstanceID      string
	ApplicationName string `hsk:"size(20)"`
	TraceEnv        tracetype.Enum
}

func (o LoginTrace) Valid() (bool, error) {
	return husk.ValidateStruct(&o)
}

func getRegistrationTrace(r Registration) LoginTrace {
	return LoginTrace{
		Allowed:         true,
		ApplicationName: r.App.Name,
		InstanceID:      r.App.InstanceID,
		IP:              r.App.IP,
		Location:        r.App.Location,
		TraceEnv:        tracetype.Register,
	}
}

func getLoginTrace(r Authentication, passed bool) LoginTrace {
	trace := tracetype.Login

	if !passed {
		trace = tracetype.Fail
	}

	return LoginTrace{
		Allowed:         passed,
		ApplicationName: r.App.Name,
		InstanceID:      r.App.InstanceID,
		IP:              r.App.IP,
		Location:        r.App.Location,
		TraceEnv:        trace,
	}
}
