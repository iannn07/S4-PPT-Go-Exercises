package Shape2Dimension

type Data struct {
	Sisi int
	Luas float64
}

func SquareArea(side int) Data {
	var area = side * side
	var dataLuas = Data{
		Sisi: side,
		Luas: float64(area),
	}
	return dataLuas
}
