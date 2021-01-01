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

//goland:noinspection GoSnakeCaseUsage
const (
	// HostType definition
	H_BROWSER  HostType = "browser"
	H_IOT      HostType = "iot"
	H_CLI      HostType = "cli"
	H_SERVER   HostType = "server"
	H_ENDPOINT HostType = "endpoint"
	H_ANY      HostType = "*"

	// Action definition
	A_LOG       Action = "log"          // From client to server: send a log message (will be store in log collection) / From server to client: send the log history
	A_READLOG   Action = "read_log"     // Request log history to the server
	A_NOTIFY    Action = "notify"       // Send a message to other endpoint
	A_REGISTER  Action = "register"     // Try to register endpoint to server
	A_REJECTED  Action = "rejected"     // Reject message from server
	A_ACCEPTED  Action = "accepted"     // Accept message from server
	A_SETUP     Action = "setup"        // Sending configuration to server
	A_ACTIVATE  Action = "activate"     // Activate a new device sent from CLI or Telegram
	A_CMD       Action = "cmd"          // Send a command to endpoint
	A_PROVISION Action = "provisioning" // Send configuration from server to IOT
	A_VALUE     Action = "value"        // Send value from IOT
	A_DISCOVER  Action = "discover"     // Send a discovery request to domain (handle by SERVER and IOT)
	A_INFORM    Action = "inform"       // From IOT or SERVER sending information about endpoint

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
	case A_LOG, A_READLOG, A_NOTIFY, A_REGISTER, A_REJECTED, A_ACCEPTED, A_SETUP, A_ACTIVATE, A_CMD, A_PROVISION, A_VALUE, A_DISCOVER, A_INFORM:
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
	// to match domain must be the same and no wildcard is allowed for domain

	if (a.Domain == comp.Domain) &&
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
	if err := json.Unmarshal(marshal, &converted); err == nil {
		return converted
	} else {
		return LogMessage{}
	}
}

func (e Endpoint) SetFromInterface(data interface{}) Endpoint {
	marshal, _ := json.Marshal(data)
	converted := Endpoint{}
	if err := json.Unmarshal(marshal, &converted); err == nil {
		return converted
	} else {
		return Endpoint{}
	}
}
