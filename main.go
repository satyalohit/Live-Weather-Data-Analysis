package main

import (
	"github.com/gin-gonic/gin"
	"github.com/satyalohit/Live-Weather-Data-Analysis/controller"
	"github.com/satyalohit/Live-Weather-Data-Analysis/weatherapi"

	"github.com/satyalohit/Live-Weather-Data-Analysis/sensor"
)

func main() {
	weatherapi.LiveWeatherApi()
	sensor.SensorApi()
	router := gin.Default()
	router.GET("/sensor", controller.GetSensorData)
	router.GET("/weather", controller.GetWeatherData)
	router.Run(":3000")
}
