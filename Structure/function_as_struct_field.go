package main

import (
	"fmt"
)

type totalSalary func(int, int, int) int

type Salary struct {
	Gross int
	Basic int
	LTA   int
	Total totalSalary
}

func main() {
	r := Salary{
		Gross: 100,
		Basic: 100,
		LTA:   100,
		Total: func(gross int, basic int, lta int) int {
			return gross + basic + lta
		},
	}
	fmt.Println(r.Total(r.Gross, r.Basic, r.LTA))
}
