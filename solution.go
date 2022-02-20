package square

import "math"

type sideCountInt uint8

const (
	SidesTriangle = 3
	SidesSquare = 4
	SidesCircle = 0
)

func CalcSquare(sideLen float64, sidesNum sideCountInt) float64 {
	switch sidesNum {
	case SidesTriangle:
		return calcTriangle(sideLen)
	case SidesCircle:
		return calcCircle(sideLen)
	default:
		return calcSquare(sideLen)
	}
}

func calcTriangle(sideLen float64) float64 {
	return math.Pow(sideLen, 2) * math.Sqrt(3.0) / float64(4)
}

func calcSquare(sideLen float64) float64 {
	return math.Pow(sideLen, 2)
}

func calcCircle(sideLen float64) float64 {
	return math.Pi * math.Pow(sideLen, 2);
}
