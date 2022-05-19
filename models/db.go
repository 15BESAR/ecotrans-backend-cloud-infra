package models

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
