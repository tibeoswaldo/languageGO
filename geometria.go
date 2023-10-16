package figuras

import "fmt"

type Geometry interface {
	Area() float64
	Perim() float64
}

func Measure(g Geometry) {
	fmt.Println("Medidas: ", g)
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimetro: ", g.Perim())
}
