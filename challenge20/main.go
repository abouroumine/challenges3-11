package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	sp, f, l := swaps(a)
	fmt.Printf("Array is sorted in %v swaps.\n", sp)
	fmt.Println("First Element:", f)
	fmt.Println("Last Element:", l)
}

func swaps(a []int32) (int, int, int) {
	numberSwaps := 0
	firstElement := -1
	lastElement := -1
	n := len(a)

	for i := 0; i < n; i++ {
		thisSwap := 0

		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				t := a[j+1]
				a[j+1] = a[j]
				a[j] = t
				thisSwap += 1
			}
		}
		if thisSwap == 0 {
			break
		}
		numberSwaps += thisSwap
	}

	firstElement = int(a[0])
	lastElement = int(a[len(a)-1])

	return numberSwaps, firstElement, lastElement
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
