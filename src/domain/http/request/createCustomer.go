package request

import "gopkg.in/go-playground/validator.v9"

type CreateCustomer struct {
	CustomerName string `json:"customer_name" validate:"required"`
}

func (r *CreateCustomer) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
