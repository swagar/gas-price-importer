package gas

import (
	"flag"
)

type Params struct {
	latitude   string
	longitude  string
	radius     string
	appKey     string
	dbURL      string
	dbUser     string
	dbPassword string
}

func ParamsFromCommndline() Params {
	var latitude *string = flag.String("latitude", "51.462", "The latitude of the desired position")
	var longitude *string = flag.String("longitude", "13.52", "The longitude of the desired position")
	var radius *string = flag.String("radius", "2.5", "The search radius of the desired position")
	var appKey *string = flag.String("appKey", "00000000-0000-0000-0000-000000000002", "App-Key for tankerkoenig")
	var dbURL *string = flag.String("dbURL", "", "URL for the InfluxDB")
	var dbUser *string = flag.String("dbUser", "", "User for the InfluxDB")
	var dbPassword *string = flag.String("dbPassword", "", "Password for the InfluxDB")
	flag.Parse()

	return Params{
		latitude:   *latitude,
		longitude:  *longitude,
		radius:     *radius,
		appKey:     *appKey,
		dbURL:      *dbURL,
		dbUser:     *dbUser,
		dbPassword: *dbPassword,
	}
}
