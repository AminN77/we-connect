package internal

import (
	"context"
	mongoPkg "github.com/AminN77/we-connect/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type Repository interface {
	Insert(fd *FinancialData) error
	InsertBatch(fd []*FinancialData) error
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

func (mr *mongoRepository) Insert(fd *FinancialData) error {
	_, err := mr.collection.InsertOne(context.Background(), fd)
	if err != nil {
		return err
	}

	return nil
}

func (mr *mongoRepository) InsertBatch(fds []*FinancialData) error {
	docs := make([]interface{}, len(fds))
	for i := 0; i < len(fds); i++ {
		docs[i] = fds[i]
	}

	_, err := mr.collection.InsertMany(context.Background(), docs)
	if err != nil {
		return err
	}

	return nil
}
