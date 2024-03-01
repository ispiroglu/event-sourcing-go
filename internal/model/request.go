package model

type CreateCustomerRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
