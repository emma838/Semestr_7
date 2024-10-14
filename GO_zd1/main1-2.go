//ZD1
// package main
// import "fmt"
// func main() {
// fmt.Println("Hello world!")
// }

// ZD2
package main

import "fmt"

func minmax(tab []int) (int, int) {
	if len(tab) < 1 {
		return 0, 0
	}

	min, max := tab[0], tab[0]

	for _, value := range tab {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}

func main() {
	numbers := []int{10, 2, 24, 13, 20}
	a, b := minmax(numbers)
	fmt.Println("Min: ", a, "Max: ", b)
}
