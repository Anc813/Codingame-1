package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"math"
)

var THOR_X, THOR_Y, ENERGY, GIANTS int

func GetDirection(x, y int) <-chan int {
	ch := make(chan int)
	go func() {
		difference := x - y
		switch {
		case difference < 0:
			ch <- -1
		case difference > 0:
			ch <- 1
		default:
			ch <- 0
		}
		close(ch)
	}()
	return ch
}

func GetDirectionLetter(a, b string, v int) string {
	switch v {
	case -1:
		return a
	case 1:
		return b
	default:
		return ""
	}
}

func Initialize(ch <-chan string) {
	fmt.Sscanf(
		<-ch,
		"%d %d",
		&THOR_X,
		&THOR_Y)
}

func Sqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func Pow(x int) int {
	return int(math.Pow(float64(x), 2.0))
}

type Position struct {
	x, y int
}

func Update(ch <-chan string) string {
	fmt.Sscanf(<-ch, "%d %d", &ENERGY, &GIANTS)

	giants := make([]Position, GIANTS)
	for i := 0; i < GIANTS; i++ {
		fmt.Sscanf(<-ch, "%d %d", &giants[i].x, &giants[i].y)
	}

	x, y := THOR_X, THOR_Y
	td, id := 9999, 0
	dc := 0

	for i, giant := range giants {
		if giant.y > y {
			dc |= 1
		} else if giant.y < y {
			dc |= 2
		}

		if giant.x > x {
			dc |= 4
		} else if giant.x < x {
			dc |= 8
		}

		dx, dy := giant.x-x, giant.y-y
		d := Sqrt(Pow(dx) + Pow(dy))

		if d < 3 {
			return "STRIKE"
		}

		if d < td {
			id = i
			td = d
		}

	}

	if dc == 15 {
		return "WAIT"
	}

	chx := GetDirection(giants[id].x, x)
	chy := GetDirection(giants[id].y, y)

	dx, dy := <-chx, <-chy
	THOR_X, THOR_Y = THOR_X+dx, THOR_Y+dy

	return GetDirectionLetter("N", "S", dy) + GetDirectionLetter("W", "E", dx)
}

func main() {
	cgreader.RunRagnarokGiants("../../input/ragnarok_giants_8.txt", true, Initialize, Update)
}
