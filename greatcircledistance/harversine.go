package greatcircledistance

import (
	"math"
	"strconv"
)

type Harverine struct{}

const earth = float64(6371)

func (gd Harverine) CalculateGreatCircleDistance(
	sourceLatitude,
	destinationLatitude,
	sourceLongitude,
	destinationLongitude string) (float64, error) {
	var kilometers float64

	sLat, err := strconv.ParseFloat(sourceLatitude, 64)
	if err != nil {
		return 0, err
	}
	dLat, err := strconv.ParseFloat(destinationLatitude, 64)
	if err != nil {
		return 0, err
	}
	sLon, err := strconv.ParseFloat(sourceLongitude, 64)
	if err != nil {
		return 0, err
	}
	dLon, err := strconv.ParseFloat(destinationLongitude, 64)
	if err != nil {
		return 0, err
	}
	sLatRad := ConvertToRadians(sLat)
	dLatRad := ConvertToRadians(dLat)
	sLonRad := ConvertToRadians(sLon)
	dLonRad := ConvertToRadians(dLon)

	deltaLat := sLatRad - dLatRad
	deltaLon := sLonRad - dLonRad

	a := math.Pow(math.Sin(deltaLat/2), 2)
	a = a + math.Cos(sLatRad)*math.Cos(dLatRad)*math.Pow(math.Sin(deltaLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	kilometers = c * earth
	return kilometers, nil
}

func ConvertToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

/*
a = sin²(Δφ/2) + cos φ1 ⋅ cos φ2 ⋅ sin²(Δλ/2)
c = 2 ⋅ atan2( √a, √(1−a) )
d = R ⋅ c
*/
