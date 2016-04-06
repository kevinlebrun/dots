package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Color struct {
	Red   int
	Green int
	Blue  int
	Alpha float32
}

func (c Color) String() string {
	return fmt.Sprintf("rgba(%d, %d, %d, %.2f)", c.Red, c.Green, c.Blue, c.Alpha)
}

func HexToColor(hex string) Color {
	red, _ := strconv.ParseInt(hex[1:3], 16, 32)
	green, _ := strconv.ParseInt(hex[3:5], 16, 32)
	blue, _ := strconv.ParseInt(hex[5:7], 16, 32)
	return Color{Red: int(red), Green: int(green), Blue: int(blue), Alpha: 1}
}

func HexesToColors(hexes string) []Color {
	var colors []Color
	for _, hex := range strings.Split(hexes, ",") {
		colors = append(colors, HexToColor(hex))
	}
	return colors
}
