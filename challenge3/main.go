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

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	fmt.Println(weirdNumber(N))
}

func weirdNumber(n int32) string {
	if n < 1 || n > 100 {
		return "Provide Proper Number!"
	}
	if n%2 != 0 {
		return "Weird"
	} else {
		if n >= 2 && n <= 5 {
			return "Not Weird"
		}
		if n >= 6 && n <= 20 {
			return "Weird"
		}
		return "Not Weird"
	}
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
