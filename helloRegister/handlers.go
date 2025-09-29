package main

import (
	"net/http"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	FirstName string
	LastName  string
	Job       string
	Age       int 
}

var users = make(map[string]User)

func register(c *gin.Context) {
	firstName := c.PostForm("firstname")
	if firstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "firtname is required"})
		return
	}
	lastName := c.PostForm("lastname")
	if lastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "lastname is required"})
		return
	}
	job := c.DefaultPostForm("job", "Unknown")

	ageStr := c.DefaultPostForm("age", "18")
	
	ageInt, err := strconv.Atoi(ageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "age should be integer"})
		return
	}

	for _, user := range users {
		if user.FirstName == firstName && user.LastName == lastName {
			c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s %s registered before", user.FirstName, user.LastName)})
			return
		}
	}

	users[firstName+lastName] = User{FirstName: firstName, LastName: lastName, Job: job, Age: ageInt}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s %s registered successfully", firstName, lastName)})
}

func hello(c *gin.Context) {
	firstName := c.Param("firstname")
	lastName := c.Param("lastname")

	if _, exists := users[firstName+lastName]; !exists {
		c.String(http.StatusNotFound, fmt.Sprintf("%s %s is not registered", firstName, lastName))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("Hello %s %s; Job: %s; Age: %d", firstName, lastName, users[firstName+lastName].Job, users[firstName+lastName].Age))
}
