package files

import (
	"context"
	"io"
	"mime/multipart"
	"os"

	"file-upload-api/domain"

	ipfsApi "github.com/ipfs/go-ipfs-api"
)

type filesUsecase struct {
	fileRepo domain.FileRepository
}

func (f filesUsecase) UploadFile(ctx context.Context, file *multipart.FileHeader) (*domain.File, error) {

	tempFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer func (){ _ = tempFile.Close() }()

	fileReader := io.Reader(tempFile)


    shell := ipfsApi.NewShellWithClient(os.Getenv("INFURA_BASE_URL"), domain.NewHttpClient(os.Getenv("INFURA_API_KEY"), os.Getenv("INFURA_API_SECRET")))

	ipfsCid, err := shell.Add(fileReader)
    if err != nil {
        return nil, err
    }


	data, err := f.fileRepo.Create(ctx, ipfsCid, file.Filename, file.Size)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f filesUsecase) ServeFile(ctx context.Context, payload *domain.GetFileByIdRequest) (io.ReadCloser, error) {

	fileObject, err := f.fileRepo.GetById(ctx, payload.FileId)
	if err != nil {
		return nil, err
	}

	fileReader, err := ipfsApi.NewLocalShell().Cat(fileObject.Cid)
	if err != nil {
		return nil, err
	}

	return fileReader, nil	
}

func (f filesUsecase) GetFileList(ctx context.Context, payload *domain.PaginatedFileListRequest) (*domain.PaginatedFileResponse, error) {

	files, err := f.fileRepo.GetFileList(ctx, int64(payload.Page), int64(payload.Limit), int64(payload.Sort))
	if err != nil {
		return nil, err
	}

	return &domain.PaginatedFileResponse{
		Meta: &domain.PaginationMeta{
			Limit: payload.Limit,
			Page: payload.Page,
		},
		Data: files,		
	}, nil
}

func New(f domain.FileRepository) domain.FilesUsecase {
	return &filesUsecase{
		fileRepo: f,
	}
}