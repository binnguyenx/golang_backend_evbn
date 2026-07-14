package handler

import (
	"net/http"

	"github.com/binnguyenx/golang_backend_evbn/internal/service"
	"github.com/gin-gonic/gin"
)

const passwordLength = 12

type GenPass interface {
}

type genPassHandler struct {
	genPassService service.GenPass
}

func NewGenPass(genPassSvc service.GenPass) GenPass {
	return &genPassHandler{
		genPassService: genPassSvc,
	}
}

func RegisterRoute(r *gin.Engine, genPassSvc service.GenPass) {
	h := &genPassHandler{genPassService: genPassSvc}
	r.GET("/genpass", h.GeneratePassword)
}

func (s *genPassHandler) GeneratePassword(c *gin.Context) {
	pass, err := s.genPassService.GeneratePassword(passwordLength)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Err"}) //we dont do err.Error because we dont want to expose the error to the client
		return
	}
	c.JSON(http.StatusOK, gin.H{"password": pass})
}
