package main

import (
	"math/rand"
	"time"
)

type ColorGenerator struct {
	colors []Color
	rand   *rand.Rand
}

func NewColorGenerator(colors []Color) *ColorGenerator {
	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return &ColorGenerator{colors, r}
}

func (c ColorGenerator) Generate() Color {
	color := c.colors[c.rand.Intn(len(c.colors))]
	color.Alpha = c.rand.Float32()
	return color
}
