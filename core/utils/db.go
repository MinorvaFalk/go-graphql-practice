package utils

import (
	"fmt"
	"log"

	"github.com/MinorvaFalk/go-graphql-practice/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func GetDatabaseGORM(l *log.Logger, env *Env) *Database {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn))

	db.AutoMigrate(models.Author{}, models.Book{})

	if err != nil {
		l.Panic(err.Error())
	}

	l.Println("Database connected")

	return &Database{db}

}
