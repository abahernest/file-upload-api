package mongodb

import (
	"github.com/Kamva/mgm/v2"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoRepository struct {
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
	}
}
