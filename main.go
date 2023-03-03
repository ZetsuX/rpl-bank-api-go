package main

import (
	"os"
	"tugas3rpl/config"
	"tugas3rpl/handler"
	"tugas3rpl/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// API Documentation : https://documenter.getpostman.com/view/25087235/2s93CUJVem

	// Setting Up Database
	database := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(database)

	// Setting Up Server
	server := gin.Default()

	// middleware CORS
	server.Use(
		middleware.CORSMiddleware(),
	)

	// Handlers
	nasabahHandler := handler.NasabahHandler{DB: database}
	noTelpHandler := handler.NoTelpHandler{DB: database}
	rekeningHandler := handler.RekeningHandler{DB: database}

	// Nasabah Routes
	server.GET("/nasabah", nasabahHandler.HandleGetNasabah)
	server.GET("/nasabah/:id", nasabahHandler.HandleGetNasabahByID)
	server.POST("/nasabah", nasabahHandler.HandleInsertNasabah)
	server.PUT("/nasabah/:id", nasabahHandler.HandleEditNasabahByID)
	server.DELETE("/nasabah/:id", nasabahHandler.HandleDeleteNasabahByID)

	// No Telp Routes
	server.GET("/notelp/nasabah/:id", noTelpHandler.HandleGetNoTelpByNasabahID)
	server.GET("/notelp", noTelpHandler.HandleGetNoTelp)
	server.GET("/notelp/:id", noTelpHandler.HandleGetNoTelpByID)
	server.POST("/notelp", noTelpHandler.HandleInsertNoTelp)
	server.DELETE("/notelp/:id", noTelpHandler.HandleDeleteNoTelpByID)
	server.PUT("/notelp/:id", noTelpHandler.HandleEditNoTelpByID)

	// Rekening Routes
	server.GET("/rekening/nasabah/:id", rekeningHandler.HandleGetRekeningByNasabahID)
	server.GET("/rekening", rekeningHandler.HandleGetRekening)
	server.GET("/rekening/:id", rekeningHandler.HandleGetRekeningByID)
	server.POST("/rekening", rekeningHandler.HandleInsertRekening)
	server.DELETE("/rekening/:id", rekeningHandler.HandleDeleteRekeningByID)
	server.PUT("/rekening/:id", rekeningHandler.HandleEditRekeningByID)

	// Running in localhost:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run("localhost:" + port)
}
