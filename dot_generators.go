package main

import (
	"math"
	"math/rand"
	"time"
)

type DotGenerator interface {
	Generate() (int, int)
}

type DotGeneratorSettings struct {
	height int
	width  int
	radius int
	rand   *rand.Rand
}

func NewDotGeneratorSettings(width, height, radius int) DotGeneratorSettings {
	return DotGeneratorSettings{
		height: height,
		width:  width,
		radius: radius,
		rand:   rand.New(rand.NewSource(time.Now().UTC().UnixNano())),
	}
}

type SimpleGenerator struct {
	width  int
	height int
	radius int
	rand   *rand.Rand
}

func NewSimpleGenerator(settings DotGeneratorSettings) *SimpleGenerator {
	return &SimpleGenerator{settings.width, settings.height, settings.radius, settings.rand}
}

func (g SimpleGenerator) Generate() (int, int) {
	x := g.rand.Float64() * float64(g.width)
	y := g.rand.NormFloat64()*100 + float64(g.height/2)
	return Constrain(int(x), g.radius, g.width-g.radius), Constrain(int(y), g.radius, g.height-g.radius)
}

type SinusoidalGenerator struct {
	width  int
	height int
	rand   *rand.Rand
}

func NewSinusoidalGenerator(settings DotGeneratorSettings) *SinusoidalGenerator {
	return &SinusoidalGenerator{settings.width, settings.height, settings.rand}
}

func (g SinusoidalGenerator) Generate() (int, int) {
	x := g.rand.Float64() * float64(g.width)
	y := math.Sin(0.007*x+30)*float64(g.height/2) + float64(g.height/2) + g.rand.NormFloat64()*100
	return int(x), int(y)
}

type CenteredGenerator struct {
	width  int
	height int
	rand   *rand.Rand
}

func NewCenteredGenerator(settings DotGeneratorSettings) *CenteredGenerator {
	return &CenteredGenerator{settings.width, settings.height, settings.rand}
}

func (g CenteredGenerator) Generate() (int, int) {
	x := g.rand.NormFloat64()*100 + float64(g.width/2)
	y := g.rand.NormFloat64()*100 + float64(g.height/2)
	return int(x), int(y)
}

type GridGenerator struct {
	width  int
	height int
	grid   [][]bool
	radius int
	rand   *rand.Rand
}

func NewGridGenerator(settings DotGeneratorSettings) *GridGenerator {
	grid := make([][]bool, settings.width/(settings.radius*2))
	for i := range grid {
		grid[i] = make([]bool, settings.height/(settings.radius*2))
	}
	return &GridGenerator{
		settings.width,
		settings.height,
		grid,
		settings.radius,
		settings.rand,
	}
}

func (g GridGenerator) Generate() (int, int) {
	idx := g.rand.Intn(g.width / (g.radius * 2))
	idy := Constrain(int(g.rand.NormFloat64()*10)+g.height/(g.radius*2)/2, 0, g.height/(g.radius*2)-1)

	for g.grid[idx][idy] == true {
		idx = g.rand.Intn(g.width / (g.radius * 2))
		idy = Constrain(int(g.rand.NormFloat64()*10)+g.height/(g.radius*2)/2, 0, g.height/(g.radius*2)-1)
	}
	g.grid[idx][idy] = true
	return idx*g.radius*2 + g.radius, idy*g.radius*2 + g.radius
}
