package main

import (
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {
		// Is it an integer?
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}
		// Is it a float
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		// Then it is invalid
		invalid = append(invalid, k)
	}
}
