package track

import (
	"context"
	"encoding/json"
	"github.com/I-Reven/Hexagonal/src/domains/entity"
	"github.com/I-Reven/Hexagonal/src/infrastructures/repository/redis"
	redisV8 "github.com/go-redis/redis/v8"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
	"os"
	"sync"
	"time"
)

var (
	once sync.Once

	client     *redisV8.Client
	ctx        context.Context
	expiration = 10 * time.Minute
	config     = &redisV8.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       1,
	}
)

func Track() redis.Redis {
	once.Do(func() {
		client = redisV8.NewClient(config)
		ctx = context.Background()
	})

	return redis.Redis{Client: client, Ctx: ctx}
}

func CreateTrack() (string, error) {
	track := entity.Track{}
	track.SetId(gocql.TimeUUID())
	id := track.GetId().String()
	track.SetId(gocql.TimeUUID())
	err := SaveTrack(id, &track)
	return id, err
}

func SaveTrack(id string, t *entity.Track) error {
	track, err := json.Marshal(t)

	if err != nil {
		err = errors.NewNotSupported(err, "error.task-can-not-connect-to-redis")
		return err
	}

	Track().Set(id, track, expiration)
	return nil
}

func GetTrack(id string) (entity.Track, error) {
	track := entity.Track{}
	t, err := Track().Get(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-track-expired")
		return track, err
	}

	err = json.Unmarshal([]byte(t), &track)

	return track, err
}

func AddMessage(id string, message string) error {
	track, err := GetTrack(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-can-not-get-track")
		return err
	}

	track.SetMessage(message)

	return SaveTrack(id, &track)
}

func AddError(id string, error error) error {
	track, err := GetTrack(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-can-not-get-track")
		return err
	}

	track.SetError(error.Error())

	return SaveTrack(id, &track)
}

func AddData(id string, data interface{}) error {
	track, err := GetTrack(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-can-not-get-track")
		return err
	}

	track.AddData(data)

	return SaveTrack(id, &track)
}

func AddDebug(id string, message string, data ...interface{}) error {
	track, err := GetTrack(id)

	if err != nil {
		err = errors.NewNotSupported(err, "error.redis-can-not-get-track")
		return err
	}

	track.AddDebug(message, data...)

	return SaveTrack(id, &track)
}
