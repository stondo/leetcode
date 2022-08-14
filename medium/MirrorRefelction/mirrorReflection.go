package mirrorrefelction

import (
	"math"
)

const NORTH = "North"
const SOUTH = "South"
const WEST = "West"
const EAST = "East"

type Point struct {
	X, Y float64
}

func getReceptorNumFromReceptorPoint(pt *Point, p float64) int {
	receptorsPositionMap := map[string]int{
		SOUTH + EAST: 0,
		NORTH + EAST: 1,
		NORTH + WEST: 2,
	}

	var xReceptorWall, yReceptorWall string

	if math.Mod(pt.X/p, 2) == 0 {
		yReceptorWall = WEST
	} else {
		yReceptorWall = EAST
	}

	if math.Mod(pt.Y/p, 2) == 0 {
		xReceptorWall = SOUTH
	} else {
		xReceptorWall = NORTH
	}

	return receptorsPositionMap[xReceptorWall+yReceptorWall]

}

func (pt *Point) findReceptorHit(fp, fq float64) int {
	isLoopEndConditionMet := func(x, y float64) bool {
		return math.Mod(x, fp) == 0 && math.Mod(y, fp) == 0
	}

	invertedSlope := fp / fq

	for i := fp; ; i += fp {
		pt.X = invertedSlope * i
		pt.Y = i
		if isLoopEndConditionMet(pt.X, pt.Y) {
			break
		}
	}

	return getReceptorNumFromReceptorPoint(pt, fp)

}

func mirrorReflection(p int, q int) int {
	origin := Point{}

	fp, fq := float64(p), float64(q)
	return origin.findReceptorHit(fp, fq)
}
