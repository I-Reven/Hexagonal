package core

import (
	"github.com/I-Reven/Hexagonal/src/infrastructure/adapter/elasticsearch"
	"github.com/I-Reven/Hexagonal/src/infrastructure/migration/cassandra"
)

type Installer struct {
	migration cassandra.Migration
	syncIndex elasticsearch.SyncIndex
}

func (i Installer) Install() error {
	if err := i.migration.Migrate(); err != nil {
		return err
	}

	if err := i.syncIndex.Send("tracker", "tracks"); err != nil {
		return err
	}

	return nil
}
