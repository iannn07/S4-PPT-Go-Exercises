package Bola

import "math"

type volume_bola struct {
	Phi float64
	Rad int
	Vol float64
}

type luas_bola struct {
	Phi  float64
	Rad  int
	Luas float64
}

// Calculate the volume of Ball
func BalVol(R int) volume_bola {
	vol := (4.0 / 3.0) * math.Pi * math.Pow(float64(R), 3)
	return volume_bola{Phi: math.Pi, Rad: R, Vol: vol}
}

func BalLuas(R int) luas_bola {
	luas_bal := 4.0 * math.Pi * math.Pow(float64(R), 2)
	return luas_bola{Phi: math.Pi, Rad: R, Luas: luas_bal}
}
