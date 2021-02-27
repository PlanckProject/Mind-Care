package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ServiceProvider struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Contact       Contact            `bson:"contact"        json:"contact"`
	Name          string             `bson:"name" json:"name"`
	ServiceType   string             `bson:"service_type"   json:"service_type"`
	FeeRange      string             `bson:"fee_range"      json:"fee_range"`
	FeeNegotiable string             `bson:"fee_negotiable" json:"fee_negotiable"`
	Timings       string             `bson:"timings" json:"timings"`
	Online        bool               `bson:"online" json:"online"`
	Declaration   bool               `bson:"declaration" json:"-"`
	Approved      bool               `bson:"approved" json:"-"`
	Location      Location           `bson:"location" json:"-"`
	CreatedAt     time.Time          `bson:"created_at" json:"-"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"-"`
}
