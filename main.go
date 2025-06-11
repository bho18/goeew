package main

import (
	"fmt"
	"sync"

	"goeew/mmi"
)

func main() {
	m := 6.9
	d := 200.0

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		fmt.Println("Bakun-Wentworth (1997):", mmi.BakunWentworth97(m, d))
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Atkinson-Wald (2007):", mmi.AtkinsonWald07(m, d))
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Allen-Wald (2012):", mmi.AllenWald12(m, d))
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Best Estimate:", mmi.BestEstimate(m, d))
	}()

	wg.Wait()
}
