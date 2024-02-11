package database

import (
	"app/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

//var dsn string = "host=localhost user=postgres password=password dbname=instaclone port=5432 sslmode=disable"

func ConnectDb() {
	var dsn string = "host=localhost user=pgpostgres password=pgpassword dbname=instaclone port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&models.User{},
		&models.Post{},
		&models.Comment{},
		&models.PostLike{},
		&models.CommentLike{},
		&models.Friendship{})

	Database = DbInstance{
		Db: db,
	}
}
