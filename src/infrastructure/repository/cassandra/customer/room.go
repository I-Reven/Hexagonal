package customer

import (
	"github.com/I-Reven/Hexagonal/src/domain/entity"
	"github.com/I-Reven/Hexagonal/src/framework/logger"
	"github.com/I-Reven/Hexagonal/src/infrastructure/repository/cassandra"
	"github.com/gocql/gocql"
	"github.com/juju/errors"
	"github.com/mitchellh/mapstructure"
	"os"
	"time"
)

var (
	toMessages = func(i interface{}) []entity.Message {
		var messages []entity.Message
		message := entity.Message{}

		if i != nil {
			for _, d := range i.([]map[string]interface{}) {
				message.SetId(d["id"].(string))
				message.SetUserId(d["user_id"].(int64))
				message.SetContent(d["content"].(string))
				message.SetKind(int32(d["kind"].(int)))
				message.SetSeen(d["seen"].([]int64))
				message.SetDelivered(d["delivered"].([]int64))
				message.SetTimestamp(d["timestamp"].(int64))

				messages = append(messages, message)
			}
		}

		return messages
	}
	toMetaData = func(i interface{}) []entity.MetaData {
		var metaData []entity.MetaData
		metaInfo := entity.MetaData{}

		if i != nil {
			for _, d := range i.([]map[string]interface{}) {
				err := mapstructure.Decode(d, &metaInfo)
				if err == nil {
					metaData = append(metaData, metaInfo)
				}
			}
		}

		return metaData
	}
	toUsersId = func(i interface{}) []int64 {
		var usersId []int64
		var userId int64

		if i != nil {
			for _, d := range i.([]int64) {
				err := mapstructure.Decode(d, &userId)
				if err == nil {
					usersId = append(usersId, userId)
				}
			}
		}

		return usersId
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
		room.SetRoomId(m["room_id"].(int64))
		room.SetStatus(int32(m["status"].(int)))
		room.SetUsersId(toUsersId(m["users_id"]))
		room.SetMessages(toMessages(m["messages"]))
		room.SetMetaData(toMetaData(m["meta_data"]))

		return room, nil
	}

	return nil, errors.NewNotFound(errors.New("error"), "error.room-not-found")
}

func (r Room) GetByRoomId(keySpace string, roomId int64) (*entity.Room, error) {
	m := map[string]interface{}{}
	query := ` SELECT * FROM rooms WHERE room_id = ? LIMIT 1 ALLOW FILTERING`
	itr := r.cql(keySpace).Query(query, roomId).Iter()
	for itr.MapScan(m) {
		room := &entity.Room{}
		room.SetId(m["id"].(gocql.UUID))
		room.SetRoomId(m["room_id"].(int64))
		room.SetStatus(int32(m["status"].(int)))
		room.SetUsersId(toUsersId(m["users_id"]))
		room.SetMessages(toMessages(m["messages"]))
		room.SetMetaData(toMetaData(m["meta_data"]))

		return room, nil
	}

	return nil, errors.NewNotFound(errors.New("error"), "error.room-not-found")
}

func (r Room) Create(keySpace string, room *entity.Room) error {
	room.SetId(gocql.TimeUUID())
	query := `INSERT INTO rooms (
				id,
				room_id,
				status,
				users_id,
				messages,
				meta_data,
				rating,
				timestamp
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
	err := r.cql(keySpace).Query(query,
		room.GetId(),
		room.GetRoomId(),
		room.GetStatus(),
		room.GetUsersId(),
		room.GetMessages(),
		room.GetMetaData(),
		room.GetRating(),
		time.Now(),
	).Exec()
	r.close()

	return err
}

func (r Room) Update(keySpace string, id gocql.UUID, room *entity.Room) error {
	query := `UPRATE rooms SET room_id = ?, status = ?, rating = ?, timestamp = ? WHERE id = ? ALLOW FILTERING;`
	err := r.cql(keySpace).Query(query,
		room.GetRoomId(),
		room.GetStatus(),
		room.GetRating(),
		time.Now(),
		id,
	).Exec()
	r.close()

	return err
}

func (r Room) AddMessage(keySpace string, id gocql.UUID, message entity.Message) error {
	query := `UPDATE rooms SET messages = messages + ?, timestamp = ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []entity.Message{message}, time.Now(), id).Exec()
	r.close()

	return err
}

func (r Room) RemoveMessage(keySpace string, id gocql.UUID, message entity.Message) error {
	query := `UPDATE rooms SET messages = messages - ?, timestamp = ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []entity.Message{message}, time.Now(), id).Exec()
	r.close()

	return err
}

func (r Room) UpdateMessage(keySpace string, id gocql.UUID, messageOld entity.Message, messageNew entity.Message) error {
	var err error

	if err = r.RemoveMessage(keySpace, id, messageOld); err != nil {
		err = r.AddMessage(keySpace, id, messageOld)
	} else {
		err = r.AddMessage(keySpace, id, messageNew)
	}

	return err
}

func (r Room) AddUser(keySpace string, id gocql.UUID, userId int64) error {
	query := `UPDATE rooms SET users_id = users_id + ?, timestamp = ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []int64{userId}, time.Now(), id).Exec()
	r.close()

	return err
}

func (r Room) RemoveUser(keySpace string, id gocql.UUID, userId int64) error {
	query := `UPDATE rooms SET users_id = users_id - ?, timestamp = ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []int64{userId}, time.Now(), id).Exec()
	r.close()

	return err
}

func (r Room) AddMetaData(keySpace string, id gocql.UUID, metaData entity.MetaData) error {
	query := `UPDATE rooms SET meta_data = meta_data + ?, timestamp = ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []entity.MetaData{metaData}, time.Now(), id).Exec()
	r.close()

	return err
}

func (r Room) RemoveMetaData(keySpace string, id gocql.UUID, metaData entity.MetaData) error {
	query := `UPDATE rooms SET meta_data = meta_data - ?, timestamp = ? WHERE id = ?`
	err := r.cql(keySpace).Query(query, []entity.MetaData{metaData}, time.Now(), id).Exec()
	r.close()

	return err
}

func (r Room) UpdateMetaData(keySpace string, id gocql.UUID, metaDataOld entity.MetaData, metaDataNew entity.MetaData) error {
	err := r.RemoveMetaData(keySpace, id, metaDataOld)

	if err != nil {
		err = r.AddMetaData(keySpace, id, metaDataOld)
	} else {
		err = r.AddMetaData(keySpace, id, metaDataNew)
	}

	return err
}
