package main

import (
	"fmt"
	"math"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var T, a int
	fmt.Scan(&T)

	var values []int

	for i := 0; i < T; i++ {
		fmt.Scan(&a)
		values = append(values, a)
	}

	for i := 0; i < T; i++ {
		isIt := isPrime(values[i])
		if isIt {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}
	}
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	sq := math.Sqrt(float64(n))
	for i := 2; i < int(sq)+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
