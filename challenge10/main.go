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
	fmt.Println(consecutiveOnes(n))
}

func baseTenToBaseTwo(n int32) string {
	value := ""
	number := n
	for number > 0 {
		rest := number % 2
		div := number / 2
		value = strconv.Itoa(int(rest)) + value
		number = div
	}
	return value
}

func consecutiveOnes(v int32) int {
	n := baseTenToBaseTwo(v)
	maxC := 1
	counter := 1
	for i := 1; i < len(n); i++ {
		if string(n[i]) == "1" {
			counter += 1
			if maxC < counter {
				maxC = counter
			}
		} else {
			counter = 0
		}
	}
	return maxC
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
