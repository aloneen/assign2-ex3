package main

import (
	"github.com/aloneen/assign2-ex3/initializers"
	"github.com/aloneen/assign2-ex3/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
