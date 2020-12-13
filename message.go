package kite_common

import (
	"errors"
	"fmt"
	"regexp"
)

type (
	HostType string
	Action   string

	Endpoint struct {
		Domain  string
		Type    HostType
		Host    string
		Address string
		Id      string
	}

	Message struct {
		Action   Action
		Sender   Endpoint
		Receiver Endpoint
		Data     interface{}
	}
)

const (
	// HostType definition
	BROWSER  HostType = "browser"
	DEVICE   HostType = "device"
	CLI      HostType = "cli"
	SERVER   HostType = "server"
	ENDPOINT HostType = "endpoint"
	ANY      HostType = "*"

	// Action definition
	LOG      Action = "log"
	NOTIFY   Action = "notify"
	REGISTER Action = "register"
	REJECTED Action = "rejected"
	ACCEPTED Action = "accepted"
	SETUP    Action = "setup"


)

func (e Endpoint) String() string {
	return fmt.Sprintf("%s.%s.%s.%s.%s", e.Domain, e.Type, e.Host, e.Address, e.Id)
}

func (e *Endpoint) StringToEndpoint(str string) {
	splitRe := regexp.MustCompile(`\.`)
	for idx, value := range splitRe.Split(str, -1) {
		if idx == 0 {
			e.Domain = value
		}
		if idx == 1 {
			e.Type = HostType(value)
			if err := e.Type.IsValid(); err != nil {
				e.Type = ANY
			}
		}
		if idx == 2 {
			e.Host = value
		}
		if idx == 3 {
			e.Address = value
		}
		if idx == 4 {
			e.Id = value
		}
	}
	if e.Domain == "" {
		e.Domain = "*"
	}
	if e.Type == "" {
		e.Type = ANY
	}
	if e.Host == "" {
		e.Host = "*"
	}
	if e.Address == "" {
		e.Address = "*"
	}
	if e.Id == "" {
		e.Id = "*"
	}
}

func (ht HostType) IsValid() error {
	switch ht {
	case BROWSER, DEVICE, CLI, SERVER, ENDPOINT, ANY:
		return nil
	}
	return errors.New("not HostType string")
}

func (a Action) IsValid() error {
	switch a {
	case LOG, NOTIFY, REGISTER, REJECTED, ACCEPTED, SETUP:
		return nil
	}
	return errors.New(fmt.Sprintf("%s is not a valid action", a))
}



func (e Endpoint) Match(comp Endpoint) bool {
	if (e.Domain == comp.Domain || comp.Domain == "*") &&
		(e.Type == comp.Type || comp.Type == "*") &&
		(e.Address == comp.Address || comp.Address == "*") &&
		(e.Host == comp.Host || comp.Host == "*") &&
		(e.Id == comp.Id || comp.Id == "*") {
		return true
	} else {
		return false
	}
}
