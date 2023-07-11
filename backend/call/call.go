package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"proj.com/connectionDB"
	"proj.com/apis/postapis"
	"proj.com/apis/getapis"
)

func main() {
	// Database
	db, err := connectionDB.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Next()
	})

	router.GET("/data", getapis.GetAllEmployees(db))
	router.GET("/data/:empid", getapis.GetResumeOfEmployee(db))
	router.POST("/submit", postapis.PostEmployeeData(db))

	router.Run("localhost:8080")
}