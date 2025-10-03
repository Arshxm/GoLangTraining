package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        username := c.GetHeader("username")
        password := c.GetHeader("password")
        if len(username) < 4 || len(password) < 4 {
            c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
            c.Abort()
            return
        }
        for _, _ = range username {
            afterPass := make([]byte, len(password))
            //reverse the username and check if it matches the password
            for i := 0; i < len(username); i++ {
                afterPass[i] = username[len(username)-i-1]
            }
            if string(afterPass) == password {
                c.Next()
                return
            }
        }
        c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        c.Abort()
    }
}