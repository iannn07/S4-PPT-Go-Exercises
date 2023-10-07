package Tabung

import "math"

type volume_cyl struct {
	Phi    float64
	Rad    int
	Height int
	Vol    float64
}

type luas_cyl struct {
	Phi    float64
	Rad    int
	Height int
	Luas   float64
}

// Calculate the volume of Ball
func CylVol(R int, H int) volume_cyl {
	vol_cyl := math.Pi * math.Pow(float64(R), 2) * float64(H)
	return volume_cyl{Phi: math.Pi, Rad: R, Height: H, Vol: vol_cyl}
}

func CylLuas(R int, H int) luas_cyl {
	luas_bal := (2.0 * math.Pi) * float64(R) * float64((R + H))
	return luas_cyl{Phi: math.Pi, Rad: R, Luas: luas_bal}
}
