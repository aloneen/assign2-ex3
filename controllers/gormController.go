package controllers

import (
	"github.com/aloneen/assign2-ex3/initializers"
	"github.com/aloneen/assign2-ex3/models"
	"github.com/gin-gonic/gin"
)

func GetUsersGORM(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func CreateUserGORM(c *gin.Context) {
	var body struct {
		Name string
		Age  int
	}
	c.Bind(&body)

	user := models.User{Name: body.Name, Age: body.Age}

	results := initializers.DB.Create(&user)

	if results.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UpdateUserGORM(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
		Age  int
	}
	c.Bind(&body)

	var user models.User

	results := initializers.DB.First(&user, id)

	if results.Error != nil {
		c.Status(400)
		return
	}
	initializers.DB.Model(&user).Updates(models.User{Name: body.Name, Age: body.Age})

	c.JSON(200, gin.H{
		"user": user,
	})

}

func DeleteUserGORM(c *gin.Context) {
	id := c.Param("id")
	results := initializers.DB.Delete(&models.User{}, id)
	if results.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"status": "deleted",
	})
}
