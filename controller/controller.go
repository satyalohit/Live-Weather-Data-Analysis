package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/satyalohit/Live-Weather-Data-Analysis/sensor"
	"github.com/satyalohit/Live-Weather-Data-Analysis/weatherapi"
)

func GetWeatherData(c *gin.Context) {
	// var startDate, endDate string
	// startDate = c.Query("startDate")
	// endDate = c.Query("endDate")
	// baseURL := fmt.Sprintf("https://archive-api.open-meteo.com/v1/archive?latitude=43.5978&longitude=-84.7675&start_date=%s&end_date=%s&hourly=temperature_2m,relativehumidity_2m,dewpoint_2m,apparent_temperature,precipitation,rain,snowfall", startDate, endDate)

	// resp, err := http.Get(baseURL)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer resp.Body.Close()

	// var weatherData weatherapi.WeatherData
	// if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
	// 	log.Println(err)
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": weatherData,
	// })
	baseURL := "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/48858/today?unitGroup=metric&key=H39TGR9W97XMEM6NV97YSTCPW&contentType=json"

	resp, err := http.Get(baseURL)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	var weatherData weatherapi.WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
			"data": weatherData,
		})
	
}

func GetSensorData(c *gin.Context) {

	// url := "https://us-east-2.aws.data.mongodb-api.com/app/data-zibxo/endpoint/data/v1/action/find"
	// method := "POST"

	// payload := strings.NewReader(`{
	// 	"collection": "readings",
	// 	"database": "decentlab_sensors",
	// 	"dataSource": "Cnets-0"
        
	// }`)

	// client := &http.Client{}
	// req, err := http.NewRequest(method, url, payload)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// req.Header.Add("Access-Control-Request-Headers", "\"*\"")
	// req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("api-key", "Q3hzYdRm2lqsl47x61WBxk3EDH7KyypRbNG7IlBjJut9Hzhb8dcU4Y8XWFX2PHBZ")

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer res.Body.Close()

	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// var oRes sensor.Document
	// err = json.Unmarshal(body, &oRes)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// var lD []sensor.Data
	// for i := 0; i < len(oRes.Doc); i++ {
	// 	var ld sensor.Data

	// 	err = json.Unmarshal([]byte(oRes.Doc[i].LoraContent), &ld)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	lD = append(lD, ld)
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"data": lD,
	// })

	url := "https://us-east-2.aws.data.mongodb-api.com/app/data-zibxo/endpoint/data/v1/action/find"
	method := "POST"
	// filter := `{
	// 	"time": {
	// 		"$gte": "2023-10-04",
	// 		"$lte": "2023-10-05"
	// 	}
	// }`

	payload := strings.NewReader(`{
		"collection": "readings",
		"database": "decentlab_sensors",
		"dataSource": "Cnets-0",
		
	}`)

	// payload := strings.NewReader(`{
	// 		"collection": "readings",
	// 		"database": "decentlab_sensors",
	// 		"dataSource": "Cnets-0",
	// 		"time": {
	// 			  "$gte": "2023-10-04T21:00:37.728+00:00",
	// 			  "$lte": "2023-10-04T22:10:38.050+00:00" 
	// 			}
	// 		}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", "Q3hzYdRm2lqsl47x61WBxk3EDH7KyypRbNG7IlBjJut9Hzhb8dcU4Y8XWFX2PHBZ")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var oRes sensor.Document
	err = json.Unmarshal(body, &oRes)
	if err != nil {
		fmt.Println(err)
		return
	}
	var lD []sensor.Data
	for _, doc := range oRes.Doc {
		var ld sensor.Data
		err = json.Unmarshal([]byte(doc.LoraContent), &ld)
		if err != nil {
			fmt.Println(err)
			continue
		}
		lD = append(lD, ld)
}
c.JSON(http.StatusOK, gin.H{
	 	"data": lD,
})

}
