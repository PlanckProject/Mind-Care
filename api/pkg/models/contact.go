package models

type Contact struct {
	ContactPerson  string   `json:"contact_person"`
	Email          string   `json:"email"`
	ContactNumbers []string `json:"contact_numbers"`
	ContactEmails  []string `json:"contact_emails"`
	Address        `json:"address"`
	Other          []interface{} `json:"other"`
	Websites       []string      `json:"websites"`
}

type Address struct {
	StreetAddress1 string `json:"street_address_1"`
	StreetAddress2 string `json:"street_address_2"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	Landmark       string `json:"landmark"`
	ZipCode        int    `json:"zip_code"`
}
