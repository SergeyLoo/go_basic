package main

import (
	"fmt"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {
	var speedInKPH EuropeanVelocity = 120.4 * 1000 / 3600
	var speedInMPS AmericanVelocity = 130 * 1609 / 3600

	fmt.Printf("Скорость в км/ч:= %f KPH\n", speedInKPH)
	fmt.Printf("Скорость 130 миль/s, в миль/ч:= %f KPH\n", speedInMPS)

}
