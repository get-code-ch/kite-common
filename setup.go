package kite_common

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	SetupFileType string

	SetupFile struct {
		Path    string `json:"path"`
		Content []byte `json:"content"`
	}

	SetupMessage struct {
		Description string      `json:"description"`
		ApiKey      string      `json:"api_key"`
		SetupFiles  []SetupFile `json:"setup_files"`
	}
)

const (
	// SetupFileType definition
	CONFIG   SetupFileType = "config"
	TELEGRAM SetupFileType = "telegram"
	CERT_KEY SetupFileType = "cert_key"
	CERT_CRT SetupFileType = "cert_crt"
)

func (s SetupFileType) IsValid() error {
	switch s {
	case CERT_CRT, CERT_KEY, CONFIG, TELEGRAM:
		return nil
	}
	return errors.New(fmt.Sprintf("%s is not a valid setup file", s))
}

func (s SetupFile) SetFromInterface(data interface{}) SetupFile {

	marshal, _ := json.Marshal(data)
	converted := SetupFile{}
	if err := json.Unmarshal(marshal, &converted); err == nil {
		return converted
	} else {
		return SetupFile{}
	}
}

func (s SetupMessage) SetFromInterface(data interface{}) SetupMessage {

	marshal, _ := json.Marshal(data)
	converted := SetupMessage{}
	if err := json.Unmarshal(marshal, &converted); err == nil {
		return converted
	} else {
		return SetupMessage{}
	}
}
