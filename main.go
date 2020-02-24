package main

import (
	"fmt"
	"wt/gas-price-importer/gas"
)

func main() {
	var params gas.Params = gas.ParamsFromCommndline()
	fmt.Println(params)

	var stations = gas.GetStationsFromTankerKoenig(params)
	fmt.Println(stations)

	gas.StoreStatationsToInfluxDB(stations, params)
}
