package main

import (
	"log"
	"os"
	"time"

	"perfect-numbers-api/internal/handlers"
	"perfect-numbers-api/internal/middleware"
	"perfect-numbers-api/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	perfectNumberService := services.NewPerfectNumberService()

	perfectNumberHandler := handlers.NewPerfectNumberHandler(perfectNumberService)

	rateLimiter := middleware.NewRateLimiter(100*time.Millisecond, 20)

	router := gin.New()

	router.Use(middleware.LoggingMiddleware())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.RateLimitMiddleware(rateLimiter))

	v1 := router.Group("/api/v1")
	{
		v1.POST("/perfect-numbers", perfectNumberHandler.FindPerfectNumbers)
		v1.GET("/health", perfectNumberHandler.Health)
		v1.GET("/info", perfectNumberHandler.APIInfo)
	}

	router.POST("/perfect-numbers", perfectNumberHandler.FindPerfectNumbers)
	router.GET("/health", perfectNumberHandler.Health)
	router.GET("/", perfectNumberHandler.APIInfo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando na porta %s", port)
	log.Printf("Endpoints disponíveis:")
	log.Printf("  POST /api/v1/perfect-numbers - Encontrar números perfeitos")
	log.Printf("  GET  /api/v1/health - Health check")
	log.Printf("  GET  /api/v1/info - Informações da API")

	if err := router.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Erro ao iniciar servidor:", err)
	}
}
