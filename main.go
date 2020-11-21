package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB


func main() {
	r := gin.Default()
	group := r.Group("/api/v1")
	{
		group.GET("/getTODO/:id", get)
		group.POST("/addTODO", post)
		group.PUT("/updateTODO/:id", put)
		group.DELETE("/deleteTODO/:id", delete)
	}
	if err := r.Run("localhost:8080"); err != nil {
		log.Println("err, message:", err)
		return
	}
	log.Println(db)
}

func get(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H {
		"id": id,
		"status":  200,
		"message": "get ok",
	})
}

func post(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"status":  200,
		"message": "add ok",
	})
}

func put(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H {
		"id": id,
		"status": 200,
		"message": "update ok",
	})
}

func delete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H {
		"id": id,
		"status": 200,
		"message": "delete ok",	
	})
}
