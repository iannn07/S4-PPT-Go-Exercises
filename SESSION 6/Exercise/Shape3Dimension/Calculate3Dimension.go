package Shape3Dimension

import (
	"math"
)

// Kindly remember that every variable in struct must be capitalized
type Data struct {
	Phi    float64
	R      int
	H      int
	Volume float64
}

func CylVolume(r int, h int) Data {
	vol := math.Pi * math.Pow(float64(r), 2) * float64(h)
	vol_struct := Data{
		Phi:    math.Pi,
		R:      r,
		H:      h,
		Volume: vol,
	}
	return vol_struct
}
