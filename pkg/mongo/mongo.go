package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoClient(connectionCtx context.Context, opts *options.ClientOptions) *mongo.Client {
	client, err := mongo.Connect(connectionCtx, opts)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func ListMongoDatabases(cli *mongo.Client, ctx context.Context) {
	if dbs, err := cli.ListDatabases(ctx, bson.D{}); err != nil {
		log.Fatal(err)
	} else {
		for _, db := range dbs.Databases {
			log.Printf("%+v", db)
		}
	}
}
