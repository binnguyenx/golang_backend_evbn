package main

import (
	"log"

	"github.com/binnguyenx/golang_backend_evbn/internal/handler"
	"github.com/binnguyenx/golang_backend_evbn/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	db, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	if err != nil {
		log.Println("db:", err)
	}
	_ = db

	genPassSvc := service.NewGenPassService()

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	handler.RegisterRoute(r, genPassSvc)

	log.Println("server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
