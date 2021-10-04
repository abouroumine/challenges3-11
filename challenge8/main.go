package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int32
	fmt.Scan(&n)
	phoneNumbers := map[string]string{}

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	for i := int32(0); i < n; i++ {
		arrTemp := strings.Split(readLine(reader), " ")
		if len(arrTemp) > 1 {
			phoneNumbers[arrTemp[0]] = arrTemp[1]
		}
	}
	stillValid := true
	for stillValid {
		arrTemp := strings.TrimSpace(readLine(reader))
		if arrTemp == "" || arrTemp == "\n" {
			stillValid = false
		}
		if stillValid {
			v, exist := phoneNumbers[arrTemp]
			if exist {
				fmt.Printf("%s=%s\n", arrTemp, v)
			} else {
				fmt.Println("Not found")
			}
		}
	}

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
