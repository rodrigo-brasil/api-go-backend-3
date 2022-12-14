package store

import (
	"fmt"
	"v0/internal/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type StoreInterface[T any] interface {
	Read(id int) (T, error)
	Create(value T) error
	Update(value T) error
	Delete(id int) error
}

type sqlStore struct {
	db *gorm.DB
}

// SetupDB : initializing mysql database
func InitDB() *gorm.DB {
	USER := "root"
	PASS := "password"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "db"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&domain.Dentista{}, &domain.Paciente{}, &domain.Consulta{})

	return db
}
