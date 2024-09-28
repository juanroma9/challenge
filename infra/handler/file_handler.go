package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
)

// The FileConfig interface represents the methods that FileHandler needs to validate if the file
// content has the correct encoding and content type. These methods are implemented by AppSettings,
// which contains the logic for that. Additionally, the FileConfig interface will be injected into
// the constructor, making use of dependency injection.
type FileConfig interface {
	IsAllowedEncoding(arrayByte []byte) (bool, error)
	IsAllowedContentType(value string) error
}

// The UserCase interface represents the method that FileHandler needs to deliver data from the file
// to the Use Case in order to execute or apply the business rules.
// This interface will be injected into the constructor, making use of dependency injection.
type UserCase interface {
	Execute(ctx context.Context, arrayBytes []byte) error
}

// FileHandler acts as a controller for file processing.
type FileHandler struct {
	config   FileConfig
	userCase UserCase
}

// NewFileHandler is a constructor and a creational design pattern
// through it, dependencies needed by NewFileHandler are injected for correct functionality.
func NewFileHandler(config FileConfig, userCase UserCase) *FileHandler {
	return &FileHandler{
		config:   config,
		userCase: userCase}
}

func (fh *FileHandler) ProcessFile(ginCtx *gin.Context) {

	// Assign 32 megabytes of max memory to load parts files.
	if err := ginCtx.Request.ParseMultipartForm(32 << 20); err != nil {
		ginCtx.String(http.StatusInsufficientStorage, err.Error())
	}

	//variables to manipulate files from the request
	var fileMultiPart multipart.File
	var fileHeader *multipart.FileHeader
	var err error

	//Iterate over all files parts from the request
	for filePart := range ginCtx.Request.MultipartForm.File {

		fileMultiPart, fileHeader, err = ginCtx.Request.FormFile(filePart)
		if err != nil {
			ginCtx.String(http.StatusInternalServerError, "Error getting file: %s", err.Error())
			return
		}
		defer fileMultiPart.Close()

		// Validate content type
		contentType := fileHeader.Header.Get("Content-Type")
		if err := fh.config.IsAllowedContentType(contentType); err != nil {
			ginCtx.String(http.StatusUnsupportedMediaType, "Invalid content type: %s", err.Error())
			return
		}

		// Read the file content.
		file, err := fileHeader.Open()
		if err != nil {
			ginCtx.String(http.StatusInternalServerError, "Error opening file: %s", err.Error())
			return
		}
		defer file.Close()

		//The file variable has a stream of bytes that corresponds to the content
		arrayBytes, err := io.ReadAll(file)
		if err != nil {
			ginCtx.String(http.StatusInternalServerError, "Error reading file: %s", err.Error())
			return
		}

		isAllowedEncondig, err := fh.config.IsAllowedEncoding(arrayBytes)
		if err != nil {
			ginCtx.String(http.StatusInternalServerError, "Error checking file encoding: %s", err.Error())
			return
		}

		if isAllowedEncondig {
			go fh.sendToGorutine(ginCtx, arrayBytes)
		}
	}

	ginCtx.JSON(http.StatusAccepted, "Se ha esta precesando la carga del archivo")
}

func (fh *FileHandler) sendToGorutine(ginCtx *gin.Context, arrayBytes []byte) {
	if userCaseError := fh.userCase.Execute(ginCtx, arrayBytes); userCaseError != nil {
		ginCtx.String(http.StatusInternalServerError, "User case execution error: %s", userCaseError.Error())
		return
	}
}
