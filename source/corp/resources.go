package corp

// IDENTIFICA OS TIPOS DE BOT APOS CARREGA-LOS

import (
	"proletor/corp/job/message"
	database "proletor/db"
)

type Resources struct {
	db *database.Database
}

func (r *Resources) load() {
	var jobs message.Console

	r.db.Connection.Where(&message.Console{}).Find(&jobs)
}

func NewResources(db *database.Database) *Resources {
	return &Resources{db}
}
