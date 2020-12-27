package kite_common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Collection string
	IcRef      string

	LogMessage struct {
		Id      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
		Time    time.Time          `bson:"time,omitempty" json:"time,omitempty"`
		Address string             `bson:"address,omitempty" json:"address,omitempty"`
		Message string             `bson:"message,omitempty" json:"message,omitempty"`
	}

	AddressAuth struct {
		Id             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
		ApiKey         string             `bson:"api_key,omitempty" json:"api_key,omitempty"`
		Name           string             `bson:"name,omitempty" json:"name,omitempty"`
		Enabled        bool               `bson:"enabled" json:"enabled,omitempty"`
		ActivationCode string             `bson:"activation_code,omitempty" json:"activation_code,omitempty"`
	}

	IC struct {
		Address     string `bson:"address" json:"address"`
		Type        IcRef  `bson:"type" json:"type"`
		Name        string `bson:"name,omitempty" json:"name,omitempty"`
		Description string `bson:"description,omitempty" json:"description,omitempty"`
	}

	Endpoint struct {
		Id         primitive.ObjectID     `bson:"_id,omitempty" json:"id,omitempty"`
		Name       string                 `bson:"name,omitempty" json:"name,omitempty"`
		IC         IC                     `bson:"ic,omitempty" json:"ic,omitempty"`
		Address    Address                `bson:"address" json:"address"`
		Attributes map[string]interface{} `bson:"attributes,omitempty" json:"attributes,omitempty"`
	}
)

const (
	// Mongo collections
	C_LOG         Collection = "log"
	C_ADDRESSAUTH Collection = "address_auth"
	C_ENDPOINT    Collection = "endpoint"

	// IcRef definition
	I_MCP23008 IcRef = "mcp23008"
	I_ADS1115  IcRef = "ads1115"
)
