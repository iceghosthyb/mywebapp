package handlers

import "github.com/gin-gonic/gin"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{"1", "Alice", "alice@example.com"},
}

func GetUsers(c *gin.Context) {
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	users = append(users, newUser)
	c.JSON(201, newUser)
}
