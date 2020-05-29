package core

import "github.com/I-Reven/Hexagonal/src/infrastructures/migration/cassandra"

func Install() error {
	return cassandra.Migrate()
}
