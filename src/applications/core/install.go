package core

import "github.com/I-Reven/Hexagonal/src/infrastructures/migration/cassandra"

type Installer struct {
	Migration cassandra.Migration
}

func (i Installer) Install() error {
	return i.Migration.Migrate()
}
