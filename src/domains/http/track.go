package http

import "gopkg.in/go-playground/validator.v9"

type Track struct {
	TrackId string `json:"trackId" validate:"required"`
}

func (r *Track) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
