package kite_common

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	HostType string
	Action   string
	Command  string

	Address struct {
		Domain  string
		Type    HostType
		Host    string
		Address string
		Id      string
	}

	Message struct {
		Action   Action
		Sender   Address
		Receiver Address
		Data     interface{}
	}
)

const (
	// HostType definition
	H_BROWSER  HostType = "browser"
	H_IOT      HostType = "iot"
	H_CLI      HostType = "cli"
	H_SERVER   HostType = "server"
	H_ENDPOINT HostType = "endpoint"
	H_ANY      HostType = "*"

	// Action definition
	A_LOG       Action = "log"
	A_READLOG   Action = "read_log"
	A_NOTIFY    Action = "notify"
	A_REGISTER  Action = "register"
	A_REJECTED  Action = "rejected"
	A_ACCEPTED  Action = "accepted"
	A_SETUP     Action = "setup"
	A_ACTIVATE  Action = "activate"
	A_CMD       Action = "cmd"
	A_PROVISION Action = "provisioning"

	// Command definition
	CMD_PUSH    Command = "push"
	CMD_READ    Command = "read"
	CMD_SET     Command = "set"
	CMD_REVERSE Command = "reverse"
)


func (ht HostType) IsValid() error {
	switch ht {
	case H_BROWSER, H_IOT, H_CLI, H_SERVER, H_ENDPOINT, H_ANY:
		return nil
	}
	return errors.New("not HostType string")
}

func (ht HostType) String() string {
	return string(ht)
}

func (a Action) IsValid() error {
	switch a {
	case A_LOG, A_READLOG, A_NOTIFY, A_REGISTER, A_REJECTED, A_ACCEPTED, A_SETUP, A_ACTIVATE, A_CMD, A_PROVISION:
		return nil
	}
	return errors.New(fmt.Sprintf("%s is not a valid action", a))
}

func (c Command) IsValid() error {
	switch c {
	case CMD_PUSH, CMD_READ, CMD_REVERSE, CMD_SET:
		return nil
	}
	return errors.New(fmt.Sprintf("%s is not a valid command", c))
}

func (a Address) Match(comp Address) bool {
	if (a.Domain == comp.Domain || comp.Domain == "*") &&
		(a.Type == comp.Type || comp.Type == "*") &&
		(a.Address == comp.Address || comp.Address == "*") &&
		(a.Host == comp.Host || comp.Host == "*") &&
		(a.Id == comp.Id || comp.Id == "*") {
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

func (e Endpoint) SetFromInterface(data interface{}) Endpoint {

	marshal, _ := json.Marshal(data)
	converted := Endpoint{}
	json.Unmarshal(marshal, &converted)
	return converted
}
