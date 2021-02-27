package models

type Contact struct {
	ContactPerson string        `bson:"contact_person" json:"contact_person"`
	Email         string        `bson:"email" json:"email"`
	Number        string        `bson:"number" json:"number"`
	Address       Address       `bson:"address" json:"address"`
	Other         []interface{} `bson:"other" json:"other"`
	Website       string        `bson:"website" json:"website"`
}

type Address struct {
	StreetAddress1 string    `bson:"street_address_1" json:"street_address_1"`
	StreetAddress2 string    `bson:"street_address_2" json:"street_address_2"`
	City           string    `bson:"city" json:"city"`
	State          string    `bson:"state" json:"state"`
	Country        string    `bson:"country" json:"country"`
	Landmark       string    `bson:"landmark" json:"landmark"`
	ZipCode        string    `bson:"zip_code" json:"zip_code"`
	Coordinates    []float64 `bson:"coordinates,omitempty" json:"coordinates,omitempty"`
}
