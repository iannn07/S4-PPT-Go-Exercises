package Limas

import "math"

type volume_limas struct {
	Sisi_alas int
	Tinggi    int
	Vol       float64
}

type luas_limas struct {
	Sisi_alas  int
	Sisi_tegak int
	Tinggi     int
	Luas       float64
}

func Limas_vol(Sisi int, Tinggi int) volume_limas {
	vol_lim := (1.0 / 3.0) * math.Pow(float64(Sisi), 2) * float64(Tinggi)
	return volume_limas{Sisi_alas: Sisi, Tinggi: Tinggi, Vol: vol_lim}
}

func Limas_luas(Sisi int, Tinggi int) luas_limas {
	segitiga := Sisi / 2
	sisi_tegak := math.Sqrt(math.Pow(float64(segitiga), 2) + math.Pow(float64(Tinggi), 2))
	luas_lim := math.Pow(float64(Sisi), 2) + (4 * sisi_tegak)
	return luas_limas{Sisi_alas: Sisi, Sisi_tegak: int(sisi_tegak), Tinggi: Tinggi, Luas: luas_lim}
}
