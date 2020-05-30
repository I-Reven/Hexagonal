package cassandra

import "github.com/gocql/gocql"

type Cassandra interface {
	InitSession() *gocql.Session
	ClearSession(session *gocql.Session)
}
