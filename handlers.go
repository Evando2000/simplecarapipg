package main

import (
	"sca-pg/models"
	"strconv"

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

func createCarHandler(c *gin.Context) {
	newReq, err := reqValidator(c)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	carModel := models.Car{
		Model: newReq.Model,
		Color: newReq.Color,
		Brand: newReq.Brand,
	}

	newCar, err := carModel.CreateCar(models.DB)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}
	c.JSON(200, gin.H{"response": newCar})
}

func getCarHandler(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	carModel := models.Car{}
	getCar, err := carModel.GetCarByID(models.DB, uid)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}
	c.JSON(200, gin.H{"response": getCar})
}

func updateCarHandler(c *gin.Context) {
	newReq, err := reqValidator(c)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	carModel := models.Car{
		Model: newReq.Model,
		Color: newReq.Color,
		Brand: newReq.Brand,
	}

	updatedCar, err := carModel.UpdateCar(models.DB, uid)
	if err != nil {
		c.JSON(404, gin.H{"response": err.Error()})
		return
	}

	c.JSON(200, gin.H{"response": updatedCar})
}

func deleteCarHandler(c *gin.Context) {
	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}

	carModel := models.Car{}
	cars, err := carModel.DeleteCar(models.DB, uid)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}
	c.JSON(200, gin.H{"response": cars})
}

func deleteAllCarHandler(c *gin.Context) {
	carModel := models.Car{}
	cars, err := carModel.DeleteAllCars(models.DB)
	if err != nil {
		c.JSON(400, gin.H{"response": err.Error()})
		return
	}
	c.JSON(200, gin.H{"response": cars})
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
