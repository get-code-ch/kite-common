package kite_common

import (
	"encoding/json"
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
	H_BROWSER  HostType = "browser"
	H_DEVICE   HostType = "device"
	H_CLI      HostType = "cli"
	H_SERVER   HostType = "server"
	H_ENDPOINT HostType = "endpoint"
	H_ANY      HostType = "*"

	// Action definition
	A_LOG      Action = "log"
	A_READLOG  Action = "read_log"
	A_NOTIFY   Action = "notify"
	A_REGISTER Action = "register"
	A_REJECTED Action = "rejected"
	A_ACCEPTED Action = "accepted"
	A_SETUP    Action = "setup"
	A_ACTIVATE Action = "activate"
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
				e.Type = H_ANY
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
		e.Type = H_ANY
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
	case H_BROWSER, H_DEVICE, H_CLI, H_SERVER, H_ENDPOINT, H_ANY:
		return nil
	}
	return errors.New("not HostType string")
}

func (a Action) IsValid() error {
	switch a {
	case A_LOG, A_READLOG, A_NOTIFY, A_REGISTER, A_REJECTED, A_ACCEPTED, A_SETUP, A_ACTIVATE:
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

func (lm LogMessage) SetFromInterface(data interface{}) LogMessage {

	marshal, _ := json.Marshal(data)
	converted := LogMessage{}
	json.Unmarshal(marshal, &converted)
	return converted
}
