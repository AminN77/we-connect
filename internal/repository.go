package internal

import (
	"context"
	mongoPkg "github.com/AminN77/we-connect/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type Repository interface {
	Insert(fd *FinancialData) error
	InsertBatch(fd []*FinancialData) error
	Get(q *Query) ([]*FinancialData, error)
}

type mongoRepository struct {
	cli        *mongo.Client
	collection *mongo.Collection
}

func NewMongoRepository() Repository {
	cli := mongoPkg.NewMongoClient(context.Background(),
		options.Client().ApplyURI(os.Getenv("MONGO_URL")))

	collection := cli.Database(os.Getenv("MONGO_DB_NAME")).
		Collection(os.Getenv("MONGO_COLLECTION_NAME"))

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

func (mr *mongoRepository) Get(q *Query) ([]*FinancialData, error) {
	var res []*FinancialData

	opts, filter, err := buildQuery(q)
	if err != nil {
		return nil, err
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Execute the query to retrieve all documents
	cursor, err := mr.collection.Find(ctx, filter, opts)

	if err != nil {
		log.Fatal(err)
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal("could not close cursor, err", err.Error())
		}
	}(cursor, ctx)

	for cursor.Next(ctx) {
		var temp *FinancialData
		err := cursor.Decode(&temp)
		if err != nil {
			return nil, err
		}

		res = append(res, temp)
	}

	// Check for errors from cursor.Err()
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
