package haversine

import (
	"math"
)

const (
	earthDiameter = 2 * 6378100
	pi180         = math.Pi / 180
)

var (
	degLatM = Distance(0, 0, 1, 0)
	degLonM = Distance(0, 0, 0, 1)
)

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance returns the distance in m between two points
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	la1 := lat1 * pi180
	la2 := lat2 * pi180
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lon2*pi180-lon1*pi180)
	return earthDiameter * math.Asin(math.Sqrt(h))
}

// Gives the top left and bottom right position arround
func BoundingBox(lat, lon, distM float64) (tlLat, tlLon, brLat, brLon float64) {
	iLat := 1.
	iLon := 1.
	if lat > 0 {
		iLat = -1
	}
	if lon > 0 {
		iLon = -1
	}
	latDegreeM := Distance(lat, lon, lat+iLat, lon)
	lonDegreeM := Distance(lat, lon, lat, lon+iLon)
	dLat := distM / latDegreeM
	dLon := distM / lonDegreeM
	tlLat = lat - dLat
	tlLon = lon + dLon
	brLat = lat + dLat
	brLon = lon - dLon
	return
}
