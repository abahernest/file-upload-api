package files

import (
	"context"
	"fmt"
	"io"

	"file-upload-api/domain"
	"file-upload-api/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var (
	validate = validator.New()
)

type FileHandler struct {
	FilesUsecase domain.FilesUsecase
	Logger *zap.Logger
}

func New(fileRouter, v1Router fiber.Router, fileUsecase domain.FilesUsecase) {
	handler := &FileHandler{
		FilesUsecase: fileUsecase,
	}

	handler.Logger, _ = logger.InitLogger()

	// v1 router -- "/api/v1/"
	v1Router.Post("/upload", handler.UploadFile)

	// file router -- "/api/v1/file/"
	fileRouter.Get("/:id", handler.GetFileById)
	fileRouter.Get("", handler.FetchFileList)
}

func (h *FileHandler) UploadFile( c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return domain.HandleError(c, err, 400, h.Logger)
	}

	file := form.File["file"][0] // accept only 1 file from the 'file' field

	data, err :=  h.FilesUsecase.UploadFile(context.TODO(), file)
	if err != nil {
		fmt.Println(err)
		return domain.HandleError(c, err, 500, h.Logger)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data": data,
	})
}

func (h *FileHandler) GetFileById( c *fiber.Ctx) error {

	var payload domain.GetFileByIdRequest

	if err := c.ParamsParser(&payload); err != nil {
		return domain.HandleError(c, err, 400, h.Logger)
	}

	if err := validate.Struct(payload); err != nil {
		return domain.HandleValidationError(c, err, h.Logger)
	}

	data, err := h.FilesUsecase.ServeFile(context.TODO(), &payload)
	if err != nil {
		return domain.HandleError(c, err, 500, h.Logger)
	}
	

	defer func(){ 
		err = data.Close()
		h.Logger.Error("error closing file", zap.Error(err))
		}()
	
	bytes, err := io.ReadAll(data)
	if err != nil {
		return domain.HandleError(c, err, 500, h.Logger)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data": string(bytes),
	})
}

func (h *FileHandler) FetchFileList( c *fiber.Ctx) error {

	var payload domain.PaginatedFileListRequest

	if err := c.QueryParser(&payload); err != nil {
		return domain.HandleError(c, err, 400, h.Logger)
	}

	if err := validate.Struct(payload); err != nil {
		return domain.HandleValidationError(c, err, h.Logger)
	}

	// set default values
	if payload.Limit ==0 {
		payload.Limit = 10
	}

	if payload.Page == 0 {
		payload.Page = 1
	}

	if payload.Sort == 0 {
		payload.Sort = -1
	}

	data, err := h.FilesUsecase.GetFileList(context.TODO(), &payload)
	if err != nil {
		return domain.HandleError(c, err, 500, h.Logger)
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data": data,
	})
}