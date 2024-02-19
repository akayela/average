// average calculates the average of several numbers.
package main

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	sample_count := float64(len(numbers))
	average := sum / float64(sample_count)

	fmt.Printf("Average: %0.2f\n", average)
}
