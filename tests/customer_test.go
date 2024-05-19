package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/mathesukkj/goloans-api/models"
	"github.com/mathesukkj/goloans-api/routes"
)

func TestCheckAvailableCustomerLoans(t *testing.T) {
	app := routes.NewRouter()

	tests := []struct {
		name     string
		customer models.Customer
		expected models.CustomerLoansResponse
	}{
		{
			name: "Customer with income <= 3000",
			customer: models.Customer{
				Age:      25,
				Cpf:      "123.456.789-00",
				Name:     "Test User",
				Income:   3000.00,
				Location: "Rio de Janeiro",
			},
			expected: models.CustomerLoansResponse{
				Customer: "Test User",
				Loans:    []models.Loan{models.PERSONAL_LOAN, models.GUARANTEED_LOAN},
			},
		},
		{
			name: "Customer with income between 3000 and 5000, age < 30, and location SP",
			customer: models.Customer{
				Age:      28,
				Cpf:      "123.456.789-00",
				Name:     "Test User",
				Income:   4000.00,
				Location: "SP",
			},
			expected: models.CustomerLoansResponse{
				Customer: "Test User",
				Loans:    []models.Loan{models.PERSONAL_LOAN, models.GUARANTEED_LOAN},
			},
		},
		{
			name: "Customer with income >= 5000",
			customer: models.Customer{
				Age:      40,
				Cpf:      "123.456.789-00",
				Name:     "Test User",
				Income:   5000.00,
				Location: "SP",
			},
			expected: models.CustomerLoansResponse{
				Customer: "Test User",
				Loans:    []models.Loan{models.CONSIGNMENT_LOAN},
			},
		},
		{
			name: "Customer with income between 3000 and 5000, age >= 30, and location SP",
			customer: models.Customer{
				Age:      30,
				Cpf:      "123.456.789-00",
				Name:     "Test User",
				Income:   4000.00,
				Location: "SP",
			},
			expected: models.CustomerLoansResponse{
				Customer: "Test User",
				Loans:    []models.Loan{},
			},
		},
		{
			name: "Customer with income between 3000 and 5000, age < 30, and location not SP",
			customer: models.Customer{
				Age:      25,
				Cpf:      "123.456.789-00",
				Name:     "Test User",
				Income:   4000.00,
				Location: "Rio de Janeiro",
			},
			expected: models.CustomerLoansResponse{
				Customer: "Test User",
				Loans:    []models.Loan{},
			},
		},
		{
			name: "Customer with income < 3000, age < 30, and location not SP",
			customer: models.Customer{
				Age:      25,
				Cpf:      "123.456.789-00",
				Name:     "Test User",
				Income:   2500.00,
				Location: "Rio de Janeiro",
			},
			expected: models.CustomerLoansResponse{
				Customer: "Test User",
				Loans:    []models.Loan{models.PERSONAL_LOAN, models.GUARANTEED_LOAN},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body bytes.Buffer
			if err := json.NewEncoder(&body).Encode(tt.customer); err != nil {
				t.Fatalf("error while encoding req body")
			}

			req := httptest.NewRequest(http.MethodPost, "/customer-loans", &body)
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("error while doing the request test")
			}

			defer resp.Body.Close()

			expected, err := json.Marshal(tt.expected)
			if err != nil {
				t.Fatalf("error while marshaling expected struct")
			}

			got, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("error while reading resp body")
			}

			if !reflect.DeepEqual(got, expected) {
				t.Errorf("got %v, expected %v", got, tt.expected)
			}
		})
	}

}
