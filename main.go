package main

import (
	"log"
	"os"

	"sca-pg/models"

	"github.com/gin-gonic/gin"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// var carList []models.Car

func prepare() {
	uploadDir := "./files"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, 0755)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	models.ConnectDatabase()

	newCar := models.Car{
		Model: "SUV",
		Color: "Black",
		Brand: "newBrand",
	}

	carCreated, err := newCar.CreateCar(models.DB)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(carCreated)
}

func main() {
	prepare()
	router := gin.Default()
	router.SetTrustedProxies([]string{"localhost", "herokuapp.com"})

	carGroupRouter := router.Group("/car")
	{
		carGroupRouter.GET("/", getAllCarsHandler)
		carGroupRouter.POST("/", createCarHandler)
		carGroupRouter.DELETE("/", deleteAllCarHandler)
		carGroupRouter.GET("/:id", getCarHandler)
		carGroupRouter.PUT("/:id", updateCarHandler)
		carGroupRouter.DELETE("/:id", deleteCarHandler)

	}

	fileGroupRouter := router.Group("/file")
	{
		fileGroupRouter.POST("/", uploadFileHandler)
	}

	port := getEnv("PORT", "9093")
	router.Run(":" + port)
}
