package company

import database "db"

type Boss struct {
	name string
}

func (b *Boss) work() {
	db := database.Connect()
	resources := NewResources(db)

	resources.load()
}

func NewBoss() *Boss {
	return &Boss{"Mautor"}
}
