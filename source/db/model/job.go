package model

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	function string

	Worker int
}
