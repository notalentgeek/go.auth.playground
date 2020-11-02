package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"go.auth.playground/handlers"
	"go.auth.playground/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := ("user=" + os.Getenv("SERVICE_POSTGRES_USER") +
		" password=" + os.Getenv("SERVICE_POSTGRES_PASSWORD") +
		" dbname=" + os.Getenv("SERVICE_POSTGRES_DB") +
		" host=" + os.Getenv("SERVICE_PGHOST") +
		" port=5432 sslmode=disable TimeZone=Asia/Jakarta")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Users{})

	usersRepository := models.UsersAdapter{DB: db}
	index := handlers.Auth{UsersRepository: usersRepository}
	auth := handlers.Auth{UsersRepository: usersRepository}

	http.HandleFunc("/home", index.Home)
	http.HandleFunc("/signup", auth.SignUp)
	http.HandleFunc("/signin", auth.SignIn)

	port := "8000"
	logrus.Infof("Server is running at 0.0.0.0:%s.", port)
	http.ListenAndServe(":"+port, nil)
}
