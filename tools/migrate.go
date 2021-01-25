package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wtnbyk-a19/youtube-manager-go/databases"
	"github.com/wtnbyk-a19/youtube-manager-go/models"
)

func main() {
	db, err := databases.Connect()
	defer db.Close()

	if err != nil {
		logrus.Fatal(err)
	}

	db.Debug().AutoMigrate(&models.User{})
	db.Debug().AutoMigrate(&models.Favorite{})
}
