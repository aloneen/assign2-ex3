package controllers

import (
	"github.com/aloneen/assign2-ex3/initializers"
	"github.com/aloneen/assign2-ex3/models"
	"github.com/gin-gonic/gin"
)

func GetUsersSQL(c *gin.Context) {
	rows, err := initializers.SQLDB.Query("SELECT id, name, age FROM users")
	if err != nil {
		c.Status(400)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Age)
		users = append(users, u)
	}

	c.JSON(200, gin.H{"users": users})
}

func CreateUserSQL(c *gin.Context) {
	var body struct {
		Name string
		Age  int
	}
	c.Bind(&body)

	var id int
	err := initializers.SQLDB.QueryRow(
		"INSERT INTO users(name, age) VALUES($1, $2) RETURNING id",
		body.Name, body.Age,
	).Scan(&id)

	if err != nil {
		c.Status(400)
		return
	}

	user := models.User{Name: body.Name, Age: body.Age}
	c.JSON(200, gin.H{"user": user})
}

func UpdateUserSQL(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name string
		Age  int
	}
	c.Bind(&body)

	_, err := initializers.SQLDB.Exec(
		"UPDATE users SET name=$1, age=$2 WHERE id=$3",
		body.Name, body.Age, id,
	)
	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"user": models.User{Name: body.Name, Age: body.Age},
	})
}

func DeleteUserSQL(c *gin.Context) {
	id := c.Param("id")

	_, err := initializers.SQLDB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"status": "deleted"})
}
