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
		KeySpace    string
		Consistency string
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

func (c *Cassandra) InitSession() *gocql.Session {
	cluster := gocql.NewCluster(c.Host)
	cluster.Port = port(c.Port)
	cluster.Keyspace = c.KeySpace
	cluster.Consistency = consistency(c.Consistency)

	session, err := cluster.CreateSession()

	if err != nil {
		err = errors.NewNotSupported(err, "error.can-not-connect-to-cassandra")
		logrus.Error(err)
	}
	return session
}

func (c *Cassandra) MackKeySpace(keySpace string) error {
	cluster := gocql.NewCluster(c.Host)
	cluster.Port = port(c.Port)
	cluster.Consistency = consistency(c.Consistency)

	session, err := cluster.CreateSession()

	if err != nil {
		err = errors.NewNotSupported(err, "error.can-not-connect-to-cassandra")
		return err
	}

	debugQuery := `create keyspace ? with replication = {'class': 'SimpleStrategy', 'replication_factor': 1};`

	return session.Query(debugQuery, keySpace).Exec()
}

func (c *Cassandra) ClearSession(session *gocql.Session) {
	session.Close()
}
