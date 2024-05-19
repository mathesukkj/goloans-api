package models

type Customer struct {
	Age      uint8  `json:"age"      validate:"required"`
	Cpf      string `json:"cpf"      validate:"required"`
	Name     string `json:"name"     validate:"required"`
	Income   int32  `json:"income"   validate:"required"`
	Location string `json:"location" validate:"required"`
}

type CustomerLoansResponse struct {
	Customer string `json:"customer"`
	Loans    []Loan `json:"loans"`
}
