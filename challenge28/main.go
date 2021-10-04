package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	var names []string

	for NItr := 0; NItr < int(N); NItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		firstName := firstMultipleInput[0]

		emailID := firstMultipleInput[1]

		for s, v := range emailID {
			if string(v) == "@" {
				if emailID[s:] == "@gmail.com" {
					names = append(names, firstName)
				}
			}
		}
	}

	if len(names) > 0 {
		sort.Strings(names)
		for _, v := range names {
			fmt.Println(v)
		}
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
