package main

import (
	"Bola"
	"Limas"
	"Tabung"
	"fmt"
	"math"
)

type BangunRuang interface {
	Volume() float64
	Luas() float64
}

type InputBall struct {
	Rad int
	Phi float64
}

type InputCylinder struct {
	Phi    float64
	Rad    int
	Height int
}

type InputLimas struct {
	SisiAlas int
	Tinggi   int
}

// Calculate the Ball
func (input InputBall) Volume() float64 {
	vol_bal := Bola.BalVol(input.Rad)
	return vol_bal.Vol
}

func (Input InputBall) Luas() float64 {
	luas_bal := Bola.BalLuas(Input.Rad)
	return luas_bal.Luas
}

// Calculate the Cylinder
func (Input InputCylinder) Volume() float64 {
	vol_cyl := Tabung.CylVol(Input.Rad, Input.Height)
	return vol_cyl.Vol
}

func (Input InputCylinder) Luas() float64 {
	luas_cyl := Tabung.CylLuas(Input.Rad, Input.Height)
	return luas_cyl.Luas
}

// Calculate the Limas
func (Input InputLimas) Volume() float64 {
	vol_cyl := Limas.Limas_vol(Input.SisiAlas, Input.Tinggi)
	return vol_cyl.Vol
}

func (Input InputLimas) Luas() float64 {
	luas_cyl := Limas.Limas_luas(Input.SisiAlas, Input.Tinggi)
	return luas_cyl.Luas
}

func main() {
	ball_data := InputBall{
		Rad: 42,
		Phi: math.Pi,
	}
	// SUCCESS
	fmt.Println("Ball")
	ball_volume := ball_data.Volume()
	fmt.Println(ball_volume)
	ball_luas := ball_data.Luas()
	fmt.Println(ball_luas)

	// SUCCESS
	cyl_data := InputCylinder{
		Phi:    math.Pi,
		Rad:    42,
		Height: 10,
	}
	fmt.Println("Cylinder")
	cyl_volume := cyl_data.Volume()
	fmt.Println(cyl_volume)
	cyl_luas := cyl_data.Luas()
	fmt.Println(cyl_luas)

	// SUCCESS
	limas_data := InputLimas{
		SisiAlas: 10,
		Tinggi:   20,
	}
	fmt.Println("Limas")
	lim_volume := limas_data.Volume()
	fmt.Println(lim_volume)
	lim_luas := limas_data.Luas()
	fmt.Println(lim_luas)
}
