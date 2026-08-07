package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	fmt.Println(convert(s, 3) == "PAHNAPLSIIGYIR")
}

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([][]byte, numRows)
	row := 0
	step := 1
	for i := range s {
		rows[row] = append(rows[row], s[i])
		switch row {
		case 0:
			step = 1
		case numRows - 1:
			step = -1
		}
		row += step
	}
	var result strings.Builder
	for i := range rows {
		result.Write(rows[i])
	}
	return result.String()
}
