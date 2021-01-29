package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ServiceProvider struct {
	ID            primitive.ObjectID `bson:"_id" json:"-"`
	Contact       Contact            `bson:"contact"        json:"contact"`
	Name          string             `bson:"name"           json:"name"`
	ServiceType   string             `bson:"service_type"   json:"service_type"`
	FeeRange      string             `bson:"fee_range"      json:"fee_range"`
	FeeNegotiable string             `bson:"fee_negotiable" json:"fee_negotiable"`
	Timings       string             `bson:"timings"        json:"timings"`
	Declaration   bool               `bson:"declaration"    json:"declaration"`
	Approved      bool               `bson:"approved"       json:"approved"`
	Location      Location           `bson:"location" json:"location"`
}
