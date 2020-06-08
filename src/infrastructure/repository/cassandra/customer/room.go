package customer

import (
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/cassandra"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
	"github.com/mitchellh/mapstructure"
	"os"
)

var (
	toMessages = func(i interface{}) []entity.Message {
		var messages []entity.Message
		message := entity.Message{}

		for _, d := range i.([]map[string]interface{}) {
			err := mapstructure.Decode(d, &message)
			if err == nil {
				messages = append(messages, message)
			}
		}

		return messages
	}
	toMetaData = func(i interface{}) []entity.MetaData {
		var metaData []entity.MetaData
		metaInfo := entity.MetaData{}

		for _, d := range i.([]map[string]interface{}) {
			err := mapstructure.Decode(d, &metaInfo)
			if err == nil {
				metaData = append(metaData, metaInfo)
			}
		}

		return metaData
	}
)

type Room struct {
	log logger.Log
}

func (r *Room) cql(keySpace string) *gocql.Session {
	cassandraConfig := cassandra.Cassandra{
		Host:        os.Getenv("ELASSANDRA_HOST"),
		Port:        os.Getenv("CASSANDRA_PORT"),
		KeySpace:    keySpace,
		Consistency: os.Getenv("CASSANDRA_CONSISTANCY_TRACKER"),
	}

	session = cassandraConfig.InitSession()

	return session
}

func (r *Room) close() {
	session.Close()
}

func (r Room) GetById(keySpace string, id gocql.UUID) (*entity.Room, error) {
	m := map[string]interface{}{}
	query := ` SELECT * FROM rooms WHERE id = ? LIMIT 1 ALLOW FILTERING`
	itr := r.cql(keySpace).Query(query, id).Iter()
	for itr.MapScan(m) {
		room := &entity.Room{}
		room.SetId(m["id"].(gocql.UUID))
		room.SetRoomId(m["roomId"].(int64))
		room.SetStatus(m["status"].(int32))
		room.SetUsersId(m["usersId"].([]int64))
		room.SetMessages(toMessages(m["messages"]))
		room.SetMetaData(toMetaData(m["metaData"]))

		return room, nil
	}

	return nil, errors.NewNotFound(errors.New("error"), "error.room-not-found")
}

func (r Room) GetByRoomId(keySpace string, roomId int64) (*entity.Room, error) {
	m := map[string]interface{}{}
	query := ` SELECT * FROM rooms WHERE roomId = ? LIMIT 1 ALLOW FILTERING`
	itr := r.cql(keySpace).Query(query, roomId).Iter()
	for itr.MapScan(m) {
		room := &entity.Room{}
		room.SetId(m["id"].(gocql.UUID))
		room.SetRoomId(m["roomId"].(int64))
		room.SetStatus(m["status"].(int32))
		room.SetUsersId(m["usersId"].([]int64))
		room.SetMessages(toMessages(m["messages"]))
		room.SetMetaData(toMetaData(m["metaData"]))

		return room, nil
	}

	return nil, errors.NewNotFound(errors.New("error"), "error.room-not-found")
}

func (r Room) Create(keySpace string, room *entity.Room) error {
	room.SetId(gocql.TimeUUID())
	query := `INSERT INTO rooms (
				id,
				roomId,
				status,
				usersId,
				messages,
				metaData,
				rating
			) VALUES (?, ?, ?, ?, ?, ?, ?);`
	err := r.cql(keySpace).Query(query,
		room.GetId(),
		room.GetRoomId(),
		room.GetStatus(),
		room.GetUsersId(),
		room.GetMessages(),
		room.GetMetaData(),
		room.GetRating(),
	).Exec()
	r.close()

	return err
}

func (r Room) Update(keySpace string, id gocql.UUID, room *entity.Room) error {
	query := `UPRATE rooms  
				SET roomId = ?, status = ?, userId = ?, messages = ?, metaData = ?, rating = ?
				WHERE id = ?;`
	err := r.cql(keySpace).Query(query,
		room.GetRoomId(),
		room.GetStatus(),
		room.GetUsersId(),
		room.GetMessages(),
		room.GetMetaData(),
		room.GetRating(),
		id,
	).Exec()
	r.close()

	return err
}

func (r Room) AddMessage(keySpace string, id gocql.UUID, message entity.Message) error {
	query := `UPDATE rooms SET messages + ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []entity.Message{message}, id).Exec()
	r.close()

	return err
}

func (r Room) RemoveMessage(keySpace string, id gocql.UUID, message entity.Message) error {
	query := `UPDATE rooms SET messages - ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []entity.Message{message}, id).Exec()
	r.close()

	return err
}

func (r Room) AddUser(keySpace string, id gocql.UUID, userId int64) error {
	query := `UPDATE rooms SET usersId + ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []int64{userId}, id).Exec()
	r.close()

	return err
}

func (r Room) RemoveUser(keySpace string, id gocql.UUID, userId int64) error {
	query := `UPDATE rooms SET usersId - ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []int64{userId}, id).Exec()
	r.close()

	return err
}

func (r Room) AddMetaData(keySpace string, id gocql.UUID, metaData entity.MetaData) error {
	query := `UPDATE rooms SET metaData + ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []entity.MetaData{metaData}, id).Exec()
	r.close()

	return err
}

func (r Room) RemoveMetaData(keySpace string, id gocql.UUID, metaData entity.MetaData) error {
	query := `UPDATE rooms SET metaData - ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []entity.MetaData{metaData}, id).Exec()
	r.close()

	return err
}
