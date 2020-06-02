package tracker

import (
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	"github.com/gocql/gocql"
)

type Track interface {
	Migrate() error
	Create(track *entity.Track) error
	GetById(id gocql.UUID) (*entity.Track, error)
	GetByTrackId(trackId gocql.UUID) (*entity.Track, error)
	Update(id gocql.UUID, message string, error string) error
	AddData(id gocql.UUID, data string) error
	RemoveData(id gocql.UUID, data string) error
	AddDebug(id gocql.UUID, debug entity.Debug) error
	RemoveDebug(id gocql.UUID, debug entity.Debug) error
}
