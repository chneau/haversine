package haversine

import "math"

const (
	earthDiameter = 2 * 6378100
	pi180         = math.Pi / 180
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
