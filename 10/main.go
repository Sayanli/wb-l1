package main

import "fmt"

func groupTemperatures(temps []float64) map[int][]float64 {
	groupTemperatures := map[int][]float64{}

	for _, v := range temps {
		vInt := int(v)
		// Шаг в 10 градусов
		step := vInt - (vInt % 10)

		groupTemperatures[step] = append(groupTemperatures[step], v)
	}

	return groupTemperatures
}

func main() {

	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	gt := groupTemperatures(temps)
	fmt.Println(gt)
}
