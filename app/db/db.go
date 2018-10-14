package db

import (
	"errors"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/skanehira/mockapi/app/common"
)

// DB connection and configuration info
type DB struct {
	*gorm.DB
	dbType  string
	logMode bool
}

// New DB
func New(dbType, file string, logMode bool) *DB {
	_, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0700)
	if err != nil {
		panic(err)
	}
	sqlite, err := gorm.Open(dbType, file)
	if err != nil {
		panic(common.NewErrConnectDB(err))
	}

	db := &DB{
		DB:     sqlite,
		dbType: dbType,
	}

	db.LogMode(logMode)
	return db
}

func (db *DB) Migration() error {
	var errs []string

	models := []interface{}{
		&Endpoint{},
		&History{},
		&Request{},
	}
	for _, model := range models {
		if err := db.AutoMigrate(model).Error; err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) != 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
}
