package domain

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/Kamva/mgm/v2"
)

type File struct {
	mgm.DefaultModel `bson:",inline"`
	Cid  string `json:"cid" bson:"cid"`
	Name string `json:"name" bson:"name"`
	Size int64 `json:"size" bson:"size"`
}


type UploadFileRequest struct {
	File  string `validate:"required" json:"file"`
}

type PaginatedFileListRequest struct {
	Limit  uint64 `json:"limit" validate:"omitempty,min=10,max=100"`
	Page  uint64 `json:"page" validate:"omitempty,min=1,max=50"`
	Sort int64 `json:"sort" validate:"omitempty,oneof=1 -1"`
}

type GetFileByIdRequest struct {
	FileId string `json:"file_id" validate:"required,mongodb"`
}

type PaginationMeta struct {
	Limit  uint64 `json:"limit"`
	Page  uint64 `json:"page"`
}

type PaginatedFileResponse struct {
	Meta  *PaginationMeta `json:"meta"`
	Data  []*File `json:"data"`
}

type FileRepository interface {
	Create(ctx context.Context, cid, filename string, fileSize int64) (*File, error)
	GetById(ctx context.Context, id string) (*File, error)
	GetFileList(ctx context.Context, page, limit, sort int64) ([]*File, error)
}

type FilesUsecase interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) (*File, error)
	ServeFile(ctx context.Context, payload *GetFileByIdRequest) (io.ReadCloser, error)
	GetFileList(ctx context.Context, pagnationQuery *PaginatedFileListRequest) (*PaginatedFileResponse, error)
}