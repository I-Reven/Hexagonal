package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructure/migration/cassandra"
)

type Installer struct {
	migration cassandra.Migration
}

func (i Installer) Install() error {
	if err := i.migration.Migrate(); err != nil {
		return err
	}

	return nil
}
