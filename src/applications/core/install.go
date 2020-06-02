package core

import "github.com/I-Reven/Hexagonal/src/infrastructure/migration/cassandra"

type Installer struct {
	Migration cassandra.Migration
}

func (i Installer) Install() error {
	return i.Migration.Migrate()
}
