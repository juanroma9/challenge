package router

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"restApi/core/domain/usecase"
	"restApi/infra/config"
	"restApi/infra/handler"
	"restApi/infra/repository"
	"time"
)

type iProcessFileHandler interface {
	ProcessFile(ginCtx *gin.Context)
}

// This function injects dependencies for the FileHandler struct.
type FileHandlerContainer struct {
	IProcessFileHandler iProcessFileHandler
}

// Constructor: Here aplay a design pattern that is creational and uses dependecy injection to load file
func NewFileHandlerContainer() (*FileHandlerContainer, error) {

	//To Load configurations from the YAML file
	appConfig := config.NewAppSettings("./config.yaml")

	// Resolve the dependency for the MySQL Repository for later injection.
	db, err := sql.Open(appConfig.AppConfig.DBMySQLConfig.DriverName, appConfig.AppConfig.DBMySQLConfig.DataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to get MySQL connection: %w", err)
	}

	// Resolve the dependency for the HttpClient for later injection
	httpClient := http.Client{Timeout: time.Duration(30) * time.Second}

	//Dependency Injector
	return &FileHandlerContainer{
		IProcessFileHandler: handler.NewFileHandler(
			appConfig,
			usecase.NewUseCaseMeliChallenge(
				appConfig,
				repository.NewMeliRepository(httpClient, appConfig.AppConfig.EndPointConfig),
				repository.NewSQLRepository(db),
			),
		),
	}, nil
}
