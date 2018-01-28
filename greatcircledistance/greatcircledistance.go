package greatcircledistance

type GreatCircleDistance interface {
	CalculateGreatCircleDistance(
		sourceLatitude,
		destinationLatitude,
		sourceLongitude,
		destinationLongitude string) (float64, error)
}
