package main

import (
	"fmt"
	"wt/gas-price-importer/gas"
)

func main() {
	var params gas.Params = gas.ParamsFromCommndline()
	fmt.Println(params)
	fmt.Println(gas.GetStationsFromTankerKoenig(params))

}
