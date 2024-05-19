package models

type Loan struct {
	Type         string `json:"type"`
	InterestRate uint8  `json:"interest_rate"`
}
