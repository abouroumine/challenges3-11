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
	var returned []int32
	var expected []int32

	for i := 0; i < 2; i++ {

		aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		if i == 0 {
			for j := 0; j < len(aTemp); j++ {
				v, _ := strconv.Atoi(aTemp[j])
				returned = append(returned, int32(v))
			}
		} else {
			for j := 0; j < len(aTemp); j++ {
				v, _ := strconv.Atoi(aTemp[j])
				expected = append(expected, int32(v))
			}
		}
	}
	if expected[2] == returned[2] {
		if expected[1] == returned[1] {
			if expected[0] == returned[0] {
				fmt.Println(0)
			} else {
				if expected[0] > returned[0] {
					fmt.Println(0)
				} else {
					fmt.Println((returned[0] - expected[0]) * 15)
				}
			}
		} else {
			if expected[1] > returned[1] {
				fmt.Println(0)
			} else {
				fmt.Println((returned[1] - expected[1]) * 500)
			}
		}
	} else {
		if expected[2] > returned[2] {
			fmt.Println(0)
		} else {
			fmt.Println(10000)
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
