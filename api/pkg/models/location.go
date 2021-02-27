package models

type Location struct {
	Type        string                 `json:"type" bson:"type"`
	Coordinates []interface{}          `json:"coordinates" bson:"coordinates"`
	Properties  map[string]interface{} `json:"properties" bson:"properties"`
}
