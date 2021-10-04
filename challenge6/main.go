package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	var T int
	var sString string
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&sString)
		printStringS(sString)
	}
}

func printStringS(s string) {
	var even, odd string

	for k, v := range s {
		if k%2 == 0 {
			even = even + string(v)
		} else {
			odd = odd + string(v)
		}
	}
	fmt.Println(even, odd)
}
