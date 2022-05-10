package main

import (
	"jimmytechnology-golang/config"
	"jimmytechnology-golang/migrations"
	"jimmytechnology-golang/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	config.InitDB()
	defer config.CloseDB()
	migrations.Migrate()

	r := gin.Default()

	r.Static("/uploads", "./uploads")

	uploadDirs := [...]string{"articles", "users"}
	for _, dir := range uploadDirs {
		os.MkdirAll("uploads/"+dir, 0755)
	}

	routes.Serve(r)

	r.Run(":" + os.Getenv("PORT"))
}
