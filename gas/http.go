package gas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
