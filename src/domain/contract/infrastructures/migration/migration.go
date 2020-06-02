package migration

type Migration interface {
	Migrate() error
}
