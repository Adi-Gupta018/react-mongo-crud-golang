package model

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Pincode   string `json:"pincode"`
}
