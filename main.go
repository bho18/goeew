package main

import (
	"fmt"
	"goEEW/mmi"
)

func main() {
	m := 6.9
	d := 200.0

	fmt.Println("Bakun-Wentworth (1997):", mmi.BakunWentworth97(m, d))
	fmt.Println("Atkinson-Wald (2007):", mmi.AtkinsonWald07(m, d))
	fmt.Println("Allen-Wald (2012):", mmi.AllenWald12(m, d))
	fmt.Println("Best Estimate:", mmi.BestEstimate(m, d))
}
