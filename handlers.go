package main

import (
	"sca-pg/models"

	"github.com/gin-gonic/gin"
)

func getAllCarsHandler(c *gin.Context) {
	carModel := models.Car{}
	carList, err := carModel.GetAllCars(models.DB)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}
	c.JSON(200, gin.H{"response": carList})
}

func uploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	path := "files/" + file.Filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": path})
}
