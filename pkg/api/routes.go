package api

import (
	"github.com/TeamStrata/strata/pkg/auth"
	"github.com/TeamStrata/strata/pkg/database"
	"github.com/google/uuid"
	"github.com/hashicorp/go-set"
	"net/http"

	"github.com/gin-gonic/gin"
)

const uuidTag = "uuid"

// Authenticate user login.
func LoginHandler(d *database.DbManager, users *set.Set[string]) gin.HandlerFunc {
	return func(c *gin.Context) {
		login := database.User{}
		err := c.ShouldBindJSON(&login)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		user, err := d.GetUserByUserName(login.Name)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		err = auth.AuthenticateUser(user.Password, login.Password)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			return
		}

		newId := addNewUUID(users)

		c.JSON(http.StatusOK, gin.H{uuidTag: newId})
	}
}

func SignUpHandler(d *database.DbManager, users *set.Set[string]) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := database.User{}
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		hash := auth.HashPassword(user.Password)
		if hash == "" {
			c.Status(http.StatusInternalServerError)
			return
		}

		err = d.InsertUser(user.Name, hash)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		newId := addNewUUID(users)
		c.JSON(http.StatusOK, gin.H{uuidTag: newId})
	}
}

// Example endpoint that returns "pong" in the response body JSON
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Add UUID to active users
func addNewUUID(users *set.Set[string]) string {
	newUserID := uuid.NewString()
	for !users.Insert(newUserID) {
		newUserID = uuid.NewString()
	}

	return newUserID
}
