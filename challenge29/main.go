package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'bitwiseAnd' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER N
 *  2. INTEGER K
 */

func bitwiseAnd(N int32, K int32) int32 {
	// Write your code here
	maxV := 0
	for i := 2; i <= int(N); i++ {
		for j := 1; j < i; j++ {
			h := i & j
			if int32(maxV) < int32(h) && int32(h) < K {
				maxV = h
			}
			if int32(maxV) == K-1 {
				return int32(maxV)
			}
		}
	}
	return int32(maxV)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		countTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		count := int32(countTemp)

		limTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		lim := int32(limTemp)

		res := bitwiseAnd(count, lim)

		fmt.Fprintf(writer, "%d\n", res)
	}

	writer.Flush()
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
