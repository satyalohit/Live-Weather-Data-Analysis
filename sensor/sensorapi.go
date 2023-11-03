package sensor

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	
)

type Response struct {
	Id          string `json:"_id"`
	LoraContent string `json:"loraContent"`
}

type Document struct {
	Doc []Response `json:"documents"`
}

func SensorApi() {

	// file, err := os.Create("output.csv")
	// if err != nil {
	// 	fmt.Println("Error creating CSV file:", err)
	// 	return
	// }
	// defer file.Close()

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

	// var oRes Document
	// err = json.Unmarshal(body, &oRes)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// var lD []Data
	// for i := 0; i < len(oRes.Doc); i++ {
	// 	var ld Data

	// 	err = json.Unmarshal([]byte(oRes.Doc[i].LoraContent), &ld)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	lD = append(lD, ld)
	// }
	// ldjson, err := json.Marshal(lD)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// json2csv.Convert(strings.NewReader(string(ldjson)), file)
	// fmt.Println(lD[0].RawData)

	url := "https://us-east-2.aws.data.mongodb-api.com/app/data-zibxo/endpoint/data/v1/action/find"
	method := "POST"

	payload := strings.NewReader(`{
			"collection": "readings",
			"database": "decentlab_sensors",
			"dataSource": "Cnets-0"
		}`)

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
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var oRes Document
	err = json.Unmarshal(body, &oRes)
	if err != nil {
		fmt.Println(err)
		return
	}

	var lData []Data

	for _, doc := range oRes.Doc {
		var ldata Data
		err = json.Unmarshal([]byte(doc.LoraContent), &ldata)
		if err != nil {
			fmt.Println(err)
			continue
		}
		lData = append(lData,ldata)

	}
	if err := ConvertToCsv(lData); err != nil {
		log.Fatal(err)
	}

}

func ConvertToCsv(data []Data) error {
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writing the header
	err = writer.Write([]string{
		"DeviceProfileName",
		"Time",
		"AirTemperature",
		"AirHumidity",
		"BarometricPressure",
		"CO2SensorTemperature",
		"Unit",
		// "DisplayName",
		// "Value",
	})
	if err != nil {
		fmt.Println("Error writing header to CSV:", err)
		return err
	}

	// Writing the data records
	for _, dataItem := range data {
		record := []string{
			dataItem.DeviceInfo.DeviceProfileName,
			dataItem.Time,
			fmt.Sprintf("%v", dataItem.Object.AirTemperature.Value),
			fmt.Sprintf("%v", dataItem.Object.AirHumidity.Value),
			fmt.Sprintf("%v", dataItem.Object.BarometricPressure.Value),
			fmt.Sprintf("%v", dataItem.Object.CO2SensorTemperature.Value),
			dataItem.Object.AirTemperature.Unit,
			// dataItem.Object.AirTemperature.DisplayName,
			// fmt.Sprintf("%v", dataItem.Object.AirTemperature.Value),
		}
		if err := writer.Write(record); err != nil {
			fmt.Println("Error writing record to CSV:", err)
			return err
		}
	}
	writer.Flush()

	return writer.Error()
}
