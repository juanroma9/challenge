package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"restApi/infra/router"
)

func main() {

	webServer := gin.Default()

	//To verify that api is online
	webServer.GET("/", func(ginCtx *gin.Context) {
		ginCtx.String(http.StatusOK, "health check")
	})

	//To validate if the configuration is allowed

	// Crear el contenedor de handlers y manejar el error
	fileHandlerContainer, err := router.NewFileHandlerContainer()
	if err != nil {
		log.Fatalf("Error creating file handler container: %v", err)
	}

	// Registrar la ruta usando el contenedor
	webServer.POST("/file", fileHandlerContainer.IProcessFileHandler.ProcessFile)

	webServer.Run()

}
