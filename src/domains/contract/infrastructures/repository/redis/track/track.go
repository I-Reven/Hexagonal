package track

import "github.com/I-Reven/Hexagonal/src/domains/entity"

type Track interface {
	CreateTrack() (string, error)
	SaveTrack(id string, track *entity.Track) error
	DeleteTrack(id string) error
	GetTrack(id string) (entity.Track, error)
	AddMessage(id string, message string) error
	AddError(id string, error error) error
	AddData(id string, data interface{}) error
	AddDebug(id string, message string, data ...interface{}) error
}
