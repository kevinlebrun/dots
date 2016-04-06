package main

import "fmt"

func Fill(color Color) string {
	return fmt.Sprintf("fill:%s", color.String())
}

func Constrain(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
