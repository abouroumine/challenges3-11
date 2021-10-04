package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != 6 {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}
	fmt.Println(computeMaxHourGlass(arr))
}

func computeHg(hg [3][3]int32) int {
	value := hg[0][0] + hg[0][1] + hg[0][2] + hg[1][1] + hg[2][0] + hg[2][1] + hg[2][2]
	return int(value)
}

func computeMaxHourGlass(hg [][]int32) int {
	maxHg := math.MinInt32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var gh = [3][3]int32{{hg[i][j], hg[i][j+1], hg[i][j+2]}, {hg[i+1][j], hg[i+1][j+1], hg[i+1][j+2]}, {hg[i+2][j], hg[i+2][j+1], hg[i+2][j+2]}}
			value := computeHg(gh)
			if value > maxHg {
				maxHg = value
			}
		}
	}
	return maxHg
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
