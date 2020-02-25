package gas

import (
	"fmt"

	client "github.com/influxdata/influxdb1-client/v2"
)

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
		point, err := client.NewPoint(station.Brand, map[string]string{
			"Name":   station.Name,
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
