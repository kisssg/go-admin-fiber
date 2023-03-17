package models

import "gorm.io/gorm"

type Entity interface {
	Count(db *gorm.DB) int64
	Take(db *gorm.DB, limit, offset int) interface{}
}
