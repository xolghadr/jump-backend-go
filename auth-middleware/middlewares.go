package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
	"strings"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		password := c.GetHeader("password")
		username := c.Request.Header.Get("username")

		if len(username) < 4 {
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("{\"message\":\"Unauthorized\"}"))
			c.Header("Content-Type", "application/json; charset=utf-8")
			return
		}
		chars := strings.Split(password, "")
		slices.Reverse(chars)
		if username != strings.Join(chars, "") {
			c.Abort()
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("{\"message\":\"Unauthorized\"}"))
			c.Header("Content-Type", "application/json; charset=utf-8")
			return
		}
		c.Next()
	}
}
