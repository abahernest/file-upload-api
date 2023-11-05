package mongodb

import (
	"context"
	"file-upload-api/domain"

	"github.com/Kamva/mgm/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type mongoFileRepository struct {
	Logger *zap.Logger
	Coll *mgm.Collection
}

func (m mongoFileRepository) Create(ctx context.Context, cid, filename string, fileSize int64) (*domain.File, error) {

	file := domain.File{Name: filename, Cid: cid, Size: fileSize}

	err := m.Coll.CreateWithCtx(ctx, &file)
	if err != nil {
		m.Logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	return &file, nil
}

func (m mongoFileRepository) GetById(ctx context.Context, id string) (*domain.File, error) {
	var file domain.File

	err := m.Coll.FindByIDWithCtx(ctx, id, &file)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		m.Logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	return &file, nil
}

func (m mongoFileRepository) GetFileList(ctx context.Context, page, limit, sort int64) ([]*domain.File, error) {
	var files []*domain.File

	pipeline := bson.A{
		bson.D{{Key: "$sort", Value: bson.M{
			"created_at": sort,
		}}},
		bson.D{{Key: "$skip", Value: (page - 1) * limit}},
		bson.D{{Key: "$limit", Value: limit}},
	}

	err := m.Coll.SimpleAggregate(&files, pipeline...)
	if err != nil {
		m.Logger.Error(err.Error(), zap.Error(err))
		return nil, err
	}

	return files, nil
}

func NewFileRepository(logger *zap.Logger) domain.FileRepository {
	return &mongoFileRepository{
		Logger: logger,
		Coll: mgm.Coll(&domain.File{}),
	}
}