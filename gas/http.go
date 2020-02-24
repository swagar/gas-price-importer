package gas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
)

func GetStationsFromTankerKoenig(params Params) []Station {
	var myClient = http.Client{Timeout: 1 * time.Second}
	var res, err = myClient.Get("https://creativecommons.tankerkoenig.de/json/list.php?lat=" + params.latitude + "&lng=" + params.longitude + "&rad=" + params.radius + "&sort=dist&type=all&apikey=" + params.appKey)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return []Station{}
	}

	defer res.Body.Close()

	var target interface{}
	err = json.NewDecoder(res.Body).Decode(&target)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return []Station{}
	}

	var result = target.(map[string]interface{})
	var ok bool = result["ok"].(bool)

	if !ok {
		var status = result["status"].(string)
		var message = result["message"].(string)

		fmt.Println("Status from TankerKönig: " + status + " => " + message)
		return []Station{}
	}

	var stations = result["stations"].([]interface{})

	var convertedStations = []Station{}

	for _, station := range stations {
		convertedStations = append(convertedStations, NewStationFromJson(station.(map[string]interface{})))
	}

	return convertedStations
}

func StoreStatationsToInfluxDB(stations []Station, params Params) error {
	connection, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     params.dbURL,
		Username: params.dbUser,
		Password: params.dbPassword,
	})

	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return err
	}

	defer connection.Close()

	bachtPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:        "gas-stations",
		RetentionPolicy: "autogen",
	})

	if err != nil {
		fmt.Println("Error creating InfluxDB bachtPoints: ", err.Error())
		return err
	}

	for _, station := range stations {
		point, err := client.NewPoint(station.Name, map[string]string{
			"Brand":  station.Brand,
			"ID":     station.ID,
			"Street": station.Street,
			"Place":  station.Place,
		},
			map[string]interface{}{
				"Diesel": station.Diesel,
				"E10":    station.E10,
				"E5":     station.E5,
			})

		if err != nil {
			fmt.Println("Error creating InfluxDB Point: ", err.Error())
			return err
		}

		bachtPoints.AddPoint(point)
	}

	err = connection.Write(bachtPoints)

	if err != nil {
		fmt.Println("Error writing batchpoint to InfluxDB: ", err.Error())
	}

	return err
}
