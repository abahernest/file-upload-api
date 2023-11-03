package mongodb

import (
	"os"
	"context"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	"file-upload-api/domain"
)

type MongoRepository struct {
	FileRepo domain.FileRepository
}

func New(l *zap.Logger) *MongoRepository {
	// connect to mongodb
	var dbName, connectionString string

	if os.Getenv("DB_NAME") != "" {
		dbName = os.Getenv("DB_NAME")
	}

	if os.Getenv("DATABASE_URL") != "" {
		connectionString = os.Getenv("DATABASE_URL")
	}

	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(connectionString), options.Client().SetMaxPoolSize(500))

	if err != nil {
		l.Error(err.Error(), zap.Error(err))
	}

	return &MongoRepository{
		FileRepo: NewFileRepository(l),
	}
}



func CustomAggregation(ctx context.Context, coll *mgm.Collection, filter bson.A, pipeline bson.A, responseObject interface{}, logger *zap.Logger) (interface{}, int64, error) {

	// seperate aggregation to count total number of documents that correspond to the Filter query
	count, err := coll.CountDocuments(ctx, bson.M{"$expr": bson.M{"$and": filter}})
	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return nil, 0, err
	}

	// main aggregation
	err = coll.SimpleAggregate(&responseObject, pipeline...)
	if err != nil {
		logger.Error(err.Error(), zap.Error(err))
		return nil, 0, err
	}

	return responseObject, count, nil
}
