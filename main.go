package main 

import ("goEEW/mmi"
		"fmt"
)



func main() {
	m := 6.9
	d := 200.0

	go fmt.Println("Equation 1: ", mmi.GetIntensity(m, d))
	go fmt.Println("Equation 2: ", mmi.GetIntensity2(m, d))
	go fmt.Println("Equation 3: ", mmi.GetIntensity3(m, d))

	var i string
	fmt.Scanln(&i)
}