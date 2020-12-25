package kite_common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Collection string

	LogMessage struct {
		Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
		Time     time.Time          `bson:"time,omitempty" json:"time,omitempty"`
		Endpoint string             `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
		Message  string             `bson:"message,omitempty" json:"message,omitempty"`
	}

	EndpointAuth struct {
		Id             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
		ApiKey         string             `bson:"api_key,omitempty" json:"api_key,omitempty"`
		Name           string             `bson:"name,omitempty" json:"name,omitempty"`
		Enabled        bool               `bson:"enabled" json:"enabled,omitempty"`
		ActivationCode string             `bson:"activation_code,omitempty" json:"activation_code,omitempty"`
	}
)

const (
	C_LOG          Collection = "log"
	C_ENDPOINTAUTH Collection = "endpoint_auth"
)
