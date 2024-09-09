package handler

import (
	"net/http"
	"strconv"
	"time"
	"training-go/go-session2/entity"

	"github.com/gin-gonic/gin"
)

var (
	users  []entity.User
	nextID int = 1
)

// create user
func CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.ID = nextID
	nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

// get user
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}
	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

// update user
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, u := range users {
		updateUser := entity.User{
			ID:        id,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: u.CreatedAt,
			UpdatedAt: time.Now(),
		}
		users[i] = updateUser
		c.JSON(http.StatusOK, updateUser)
		return

	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
}

// delete user
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})

}

// get all users
func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
