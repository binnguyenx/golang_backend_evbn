package api

import (
	"github.com/gin-gonic/gin"
)

type Engine interface {
	GeneratePassword(c *gin.Context)
}

type engine struct {
	app *gin.Engine //dependency injection
}

func NewEngine() Engine {
	return &engine{
		app: gin.Default(),
	}
}

func (e *engine) Start() error {
	return e.app.Run(":8080")
}

func (e *engine) initRoutes() {
	//genpass handler
	genPassSvc := service.NewGenPassService()
	genPassHandler := handler.NewGenPass(genPassSvc)
	
	e.app.POST("/genpass", genPassHandler.GeneratePassword)
}