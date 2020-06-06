package request

import "gopkg.in/go-playground/validator.v9"

type CreateRoom struct {
	CustomerName string `json:"customer_name" validate:"required"`
	RoomId       string `json:"room_id" validate:"required"`
	UserId       string `json:"user_id" validate:"required"`
}

func (r *CreateRoom) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
