package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Repo interface {
	UserRepo
}

type repo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewRepo() repo {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URL")))
	if err != nil {
		log.Fatal("Error : " + err.Error())
	}

	db := c.Database("MyData")

	return repo{Client: c, DB: db}
}
