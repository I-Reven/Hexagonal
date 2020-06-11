package request

import "gopkg.in/go-playground/validator.v9"

type AddMessage struct {
	CustomerName string `json:"customer_name" validate:"required"`
	RoomId       int64  `json:"room_id" validate:"required"`
	UserId       int64  `json:"user_id" validate:"required"`
	Content      string `json:"content" validate:"required"`
	Kind         int64  `json:"kind" validate:"required"`
}

func (r *AddMessage) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
