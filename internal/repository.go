package internal

import (
	"context"
	"errors"
	mongoPkg "github.com/AminN77/we-connect/pkg/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	ErrDatabase = errors.New("some error occurred on the database side")
)

// Repository interface is implementation of the famous repository pattern for decoupling the
// database from the domain
type Repository interface {
	Insert(fd *FinancialData) error

	// InsertBatch is more performant with bulk data
	InsertBatch(fd []*FinancialData) error

	Get(q *Query, ctx context.Context) ([]*FinancialData, error)
}

// mongoRepository handles the mongo functionality combined with repository pattern
type mongoRepository struct {
	cli        *mongo.Client
	collection *mongo.Collection
}

func NewMongoRepository() Repository {
	opts := options.Client()

	// In the case of heavy load connection pool with bigger size is recommended
	// opts.SetMaxPoolSize(1000)

	opts.ApplyURI(os.Getenv("MONGO_URL"))
	cli := mongoPkg.NewMongoClient(context.Background(), opts)

	collection := cli.Database(os.Getenv("MONGO_DB_NAME")).
		Collection(os.Getenv("MONGO_COLLECTION_NAME"))

	//// full text search index on seriesTitle1
	//keys := bson.M{"seriesTitle1": "text"}
	//indexModel := mongo.IndexModel{
	//	Keys:    keys,
	//	Options: options.Index(),
	//}
	//
	//_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	//if err != nil {
	//	log.Println("index could not been created on seriesTitle1, err", err.Error())
	//}

	return &mongoRepository{
		cli:        cli,
		collection: collection,
	}
}

func (mr *mongoRepository) Insert(fd *FinancialData) error {
	_, err := mr.collection.InsertOne(context.Background(), fd)
	if err != nil {
		log.Println(err)
		return ErrDatabase
	}

	return nil
}

func (mr *mongoRepository) InsertBatch(fds []*FinancialData) error {
	docs := make([]interface{}, len(fds))
	for i := 0; i < len(fds); i++ {
		docs[i] = fds[i]
	}

	opts := options.InsertMany()
	opts.SetOrdered(false)
	_, err := mr.collection.InsertMany(context.Background(), docs, opts)
	if err != nil {
		log.Println(err)
		return ErrDatabase
	}

	return nil
}

func (mr *mongoRepository) Get(q *Query, ctx context.Context) ([]*FinancialData, error) {
	var res []*FinancialData

	opts, filter, err := buildQuery(q)
	if err != nil {
		log.Println(err.Error())
		return nil, ErrDatabase
	}

	// Execute the query to retrieve all documents
	cursor, err := mr.collection.Find(ctx, filter, opts)
	if err != nil {
		log.Println(err)
		return nil, ErrDatabase
	}

	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Println("could not close cursor, err", err.Error())
		}
	}(cursor, ctx)

	for cursor.Next(ctx) {
		var temp *FinancialData
		err := cursor.Decode(&temp)
		if err != nil {
			log.Println(err)
			return nil, ErrDatabase
		}

		res = append(res, temp)
	}

	// Check for errors from cursor.Err()
	if err := cursor.Err(); err != nil {
		log.Println(err)
		return nil, ErrDatabase
	}

	return res, nil
}
