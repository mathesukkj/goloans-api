package models

type Customer struct {
	Age      uint8  `json:"age"`
	Cpf      string `json:"cpf"`
	Name     string `json:"name"`
	Income   int32  `json:"income"`
	Location string `json:"location"`
}
