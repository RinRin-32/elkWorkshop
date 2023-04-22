package main

import (
	"elkWorkshop/db"
	"elkWorkshop/models"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	dbConfig := db.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "elkWorkshop",
		Password: "elkWorkshop_secrets",
		DbName:   "elkWorkshop_db",
	}
	fmt.Println("%v", dbConfig)
	dbInstance, err := db.Init(dbConfig)
	if err != nil {
		logger.Err(err).Msg("Connection failed")
		os.Exit(1)
	}

	for i := 0; i < 20; i++ {
		post := &models.Post{
			Title: faker.Sentence(),
			Body:  faker.Paragraph(),
		}
		err = dbInstance.SavePost(post)
		if err != nil {
			log.Err(err).Msg("failed to save record")
		}
	}
}
