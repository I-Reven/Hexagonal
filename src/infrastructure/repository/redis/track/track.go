package track

import (
	"context"
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	repository "github.com/I-Reven/Hexagonal/src/infrastructure/repository/redis"
	redisV8 "github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
	"os"
	"sync"
	"time"
)

var (
	once       sync.Once
	redis      repository.Redis
	expiration = 10 * time.Minute
	config     = &redisV8.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       1,
	}
)

type Track struct{}

func (r Track) redis() *repository.Redis {
	once.Do(func() {
		redis = repository.Redis{Client: redisV8.NewClient(config), Ctx: context.Background()}
	})

	return &redis
}

func (r Track) CreateTrack() (string, error) {
	track := entity.Track{}
	track.SetId(gocql.TimeUUID())
	id := track.GetId().String()
	track.SetTrackId(track.GetId())
	err := r.SaveTrack(id, &track)
	return id, err
}

func (r Track) SaveTrack(id string, track *entity.Track) error {
	Track, err := json.Marshal(track)

	if err != nil {
		err = errors.NewNotSupported(err, "error.task-can-not-connect-to-redis")
		return err
	}

	return r.redis().Set(id, Track, expiration)
}

func (r Track) DeleteTrack(id string) error {
	return r.redis().Del(id)
}

func (r Track) GetTrack(id string) (entity.Track, error) {
	track := entity.Track{}
	Track, err := r.redis().Get(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-track-expired")
		return track, err
	}

	err = json.Unmarshal([]byte(Track), &track)

	return track, err
}

func (r Track) AddMessage(id string, message string) error {
	track, err := r.GetTrack(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-can-not-get-track")
		return err
	}

	track.SetMessage(message)

	return r.SaveTrack(id, &track)
}

func (r Track) AddError(id string, error error) error {
	track, err := r.GetTrack(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-can-not-get-track")
		return err
	}

	track.SetError(error.Error())

	return r.SaveTrack(id, &track)
}

func (r Track) AddData(id string, data interface{}) error {
	track, err := r.GetTrack(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-can-not-get-track")
		return err
	}

	track.AddData(data)

	return r.SaveTrack(id, &track)
}

func (r Track) AddDebug(id string, message string, data ...interface{}) error {
	track, err := r.GetTrack(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-can-not-get-track")
		return err
	}

	track.AddDebug(message, data...)

	return r.SaveTrack(id, &track)
}
