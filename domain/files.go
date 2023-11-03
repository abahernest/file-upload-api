package domain

import (
	"context"

	"github.com/Kamva/mgm/v2"
)

type File struct {
	mgm.DefaultModel `bson:",inline"`
	Cid  string `json:"cid" bson:"cid"`
	Name string `json:"name" bson:"name"`
	Size float64 `json:"size" bson:"size"`
}


type UploadFileRequest struct {

}

type FileRepository interface {
	Create(ctx context.Context, cid, filename string, fileSize float64) (*File, error)
	GetById(ctx context.Context, id string) (*File, error)
	GetFileList(ctx context.Context, page, limit int64) ([]*File, error)
}

type FilesUsecase interface {

}