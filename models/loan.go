package models

type Loan struct {
	Type         string `json:"type"`
	InterestRate uint8  `json:"interest_rate"`
}

var PERSONAL_LOAN = Loan{
	Type:         "PERSONAL",
	InterestRate: 4,
}

var GUARANTEED_LOAN = Loan{
	Type:         "GUARANTEED",
	InterestRate: 3,
}

var CONSIGMENT_LOAN = Loan{
	Type:         "CONSIGNMENT",
	InterestRate: 2,
}
