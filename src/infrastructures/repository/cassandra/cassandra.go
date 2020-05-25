package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"strconv"
)

type (
	Cassandra struct {
		Host        string
		Port        string
		Keyspace    string
		Consistancy string
	}
)

var (
	port = func(p string) int {
		i, err := strconv.Atoi(p)
		if err != nil {
			return 9042
		}

		return i
	}

	consistency = func(c string) gocql.Consistency {
		return gocql.ParseConsistency(c)
	}
)

func (c Cassandra) InitSession() *gocql.Session {
	cluster := gocql.NewCluster(c.Host)
	cluster.Port = port(c.Port)
	cluster.Keyspace = c.Keyspace
	cluster.Consistency = consistency(c.Consistancy)

	session, err := cluster.CreateSession()

	if err != nil {
		err = errors.NewNotSupported(err, "Can not connect to cassandra")
		logrus.Error(err)
	}
	return session
}

func ClearSession(session *gocql.Session) {
	session.Close()
}
