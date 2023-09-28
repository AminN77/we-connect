package internal

import (
	"context"
	mongoPkg "github.com/AminN77/we-connect/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Repository interface {
	Add(fd *FinancialData) error
}

type mongoRepository struct {
	cli        *mongo.Client
	collection *mongo.Collection
}

func NewMongoRepository() Repository {
	cli := mongoPkg.NewMongoClient(context.Background(),
		options.Client().ApplyURI(os.Getenv("MONGO_URL")))

	collection := cli.Database(os.Getenv("MONGO_DB_NAME")).Collection(os.Getenv("MONGO_COLLECTION_NAME"))

	return &mongoRepository{
		cli:        cli,
		collection: collection,
	}
}

func (mr *mongoRepository) Add(fd *FinancialData) error {
	_, err := mr.collection.InsertOne(context.Background(), fd)
	if err != nil {
		return err
	}

	return nil
}
