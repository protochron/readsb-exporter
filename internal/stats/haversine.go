package stats

import "math"

const earthRadiusMeters = 6371000

type LatLon struct {
	Lat float64
	Lon float64
}

func radians(d float64) float64 {
	return d * math.Pi / 180
}

func Haversine(pos1, pos2 LatLon) float64 {
	lat1 := radians(pos1.Lat)
	lon1 := radians(pos1.Lon)
	lat2 := radians(pos2.Lat)
	lon2 := radians(pos2.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(diffLon/2), 2)

	return 2 * earthRadiusMeters * math.Asin(a)
}
