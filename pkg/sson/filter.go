package sson

import "go.mongodb.org/mongo-driver/bson"

type E struct {
	Key   string
	Value interface{}
}

type D bson.D

type M map[string]interface{}
