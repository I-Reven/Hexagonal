package request

import "gopkg.in/go-playground/validator.v9"

type AddUser struct {
	CustomerName string `json:"customer_name" validate:"required"`
	RoomId       int64  `json:"room_id" validate:"required"`
	UserId       int64  `json:"user_id" validate:"required"`
}

func (r *AddUser) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
