package request

import "gopkg.in/go-playground/validator.v9"

type AddMetadata struct {
	CustomerName string `json:"customer_name" validate:"required"`
	RoomId       int64  `json:"room_id" validate:"required"`
	Key          string `json:"key" validate:"required"`
	Kind         int32  `json:"kind" validate:"required"`
	Value        string `json:"value" validate:"required"`
}

func (r *AddMetadata) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
