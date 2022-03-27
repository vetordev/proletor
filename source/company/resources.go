package company

import (
	database "db"
	"message"
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
