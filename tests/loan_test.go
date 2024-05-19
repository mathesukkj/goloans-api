package tests

import (
	"reflect"
	"testing"

	"github.com/mathesukkj/goloans-api/models"
	"github.com/mathesukkj/goloans-api/services"
)

func TestGetAvailableLoans(t *testing.T) {
	tests := []struct {
		name     string
		customer models.Customer
		expected []models.Loan
	}{
		{
			name:     "Customer with income <= 3000",
			customer: models.Customer{Income: 3000, Age: 25, Location: "Rio de Janeiro"},
			expected: []models.Loan{models.PERSONAL_LOAN, models.GUARANTEED_LOAN},
		},
		{
			name:     "Customer with income between 3000 and 5000, age < 30, and location SP",
			customer: models.Customer{Income: 4000, Age: 28, Location: "SP"},
			expected: []models.Loan{models.PERSONAL_LOAN, models.GUARANTEED_LOAN},
		},
		{
			name:     "Customer with income >= 5000",
			customer: models.Customer{Income: 5000, Age: 40, Location: "SP"},
			expected: []models.Loan{models.CONSIGMENT_LOAN},
		},
		{
			name:     "Customer with income between 3000 and 5000, age >= 30, and location SP",
			customer: models.Customer{Income: 4000, Age: 30, Location: "SP"},
			expected: []models.Loan{},
		},
		{
			name:     "Customer with income between 3000 and 5000, age < 30, and location not SP",
			customer: models.Customer{Income: 4000, Age: 25, Location: "Rio de Janeiro"},
			expected: []models.Loan{},
		},
		{
			name:     "Customer with income < 3000, age < 30, and location not SP",
			customer: models.Customer{Income: 2500, Age: 25, Location: "Rio de Janeiro"},
			expected: []models.Loan{models.PERSONAL_LOAN, models.GUARANTEED_LOAN},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := services.GetAvailableLoans(&tt.customer)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Available loans = %v, expected %v", got, tt.expected)
			}
		})
	}

}
