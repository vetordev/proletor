package model

import "gorm.io/gorm"

type Worker struct {
	gorm.Model
	Name string

	Jobs []Job `gorm:"foreignKey:Worker"`
	//owner Owner
}
