package services

import "github.com/mathesukkj/goloans-api/models"

func GetAvailableLoans(customer *models.Customer) []models.Loan {
	loans := []models.Loan{}

	if customer.Income <= 3000 {
		loans = append(loans, models.PERSONAL_LOAN)
		loans = append(loans, models.GUARANTEED_LOAN)
	}

	if customer.Income > 3000 && customer.Income < 5000 && customer.Age < 30 &&
		customer.Location == "SP" {
		loans = append(loans, models.PERSONAL_LOAN)
		loans = append(loans, models.GUARANTEED_LOAN)
	}

	if customer.Income >= 5000 {
		loans = append(loans, models.CONSIGNMENT_LOAN)
	}

	return loans
}
