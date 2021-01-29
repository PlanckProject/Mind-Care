package models

type Location struct {
	Type     string `json:"type" bson:"type"`
	Geometry struct {
		Type        string        `json:"type" bson:"type"`
		Coordinates []interface{} `json:"coordinates" bson:"coordinates"`
	} `json:"geometry" bson:"geometry"`
	Properties map[string]interface{} `json:"properties" bson:"properties"`
}
