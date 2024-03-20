package model

type Citizen struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	DOB       string `json:"dob"`
	Gender    string `json:"gender"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Pincode   string `json:"pincode"`
}
