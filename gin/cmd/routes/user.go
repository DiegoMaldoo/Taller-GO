package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users []User

func SetupUserRoutes(r *gin.Engine) {
	r.GET("/users", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		fmt.Println("Request user agent: ", userAgent)
		c.Header("x-user-agent", "gin")
		c.JSON(http.StatusOK, users)
	})
	r.POST("/users", func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"mesage": "error parsing body",
			})
			return
		}
		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"mesage": "error parsing body",
			})
			return
		}
		user.ID = len(users) + 1
		users = append(users, user)

		c.JSON(http.StatusOK, user)
	})
	r.PUT("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		for i, user := range users {
			if fmt.Sprintf("%d", user.ID) == idParam {
				var updatedUser User
				if err := c.ShouldBindJSON(&updatedUser); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
					return
				}

				users[i].Name = updatedUser.Name
				users[i].Email = updatedUser.Email
				c.JSON(http.StatusOK, users[i])
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		for i, user := range users {
			if fmt.Sprintf("%d", user.ID) == idParam {
				users = append(users[:i], users[i+1:]...)
				c.JSON(http.StatusNoContent, nil)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	})
}
