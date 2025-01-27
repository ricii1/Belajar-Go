package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Age   int    `json:"age" binding:"required"`
}

func main() {
	r := gin.Default()

	var users []User

	// Root Endpoint
	// Context adalah objek yang berisi informasi tentang request yang sedang berjalan
	// Dapat dianalogikan seperti pelayan yang mengetahuhi semua permintaan dari pelanggan
	// Pelayan tahu apa yang diminta pelanggan (request).
	// Pelayan tahu apa yang harus diberikan kembali ke pelanggan (response).
	// Jika ada masalah, pelayan bisa memberitahu pelanggan dengan sopan (error handling).
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Root Endpoint",
		})
	})

	// Get Endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Get Endpoint with Query
	r.GET("/query/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
		})
	})

	// Get Endpoint with Body
	// Note: Tidak semua bahasa mendukung get dengan body
	r.GET("/body", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name":  user.Name,
			"email": user.Email,
			"age":   user.Age,
		})
	})
	// Post Endpoint
	r.POST("/post", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, user)
		c.JSON(http.StatusOK, gin.H{
			"name":  user.Name,
			"email": user.Email,
			"age":   user.Age,
		})
	})
	// Run server in port 8080
	r.Run(":8080")
}
