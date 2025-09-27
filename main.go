package main

import (
	"github.com/aloneen/assign2-ex3/controllers"
	"github.com/aloneen/assign2-ex3/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()

	// GORM
	r.GET("/gorm/users", controllers.GetUsersGORM)
	r.POST("/gorm/users", controllers.CreateUserGORM)
	r.PUT("/gorm/users/:id", controllers.UpdateUserGORM)
	r.DELETE("/gorm/users/:id", controllers.DeleteUserGORM)

	// SQL
	r.GET("/sql/users", controllers.GetUsersSQL)
	r.POST("/sql/users", controllers.CreateUserSQL)
	r.PUT("/sql/users/:id", controllers.UpdateUserSQL)
	r.DELETE("/sql/users/:id", controllers.DeleteUserSQL)

	r.Run()
}
