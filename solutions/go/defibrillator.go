package main

import (
	"fmt"
	"github.com/glendc/cgreader"
	"math"
	"strings"
)

func ParseFloat(in string) (out float64) {
	fmt.Sscanf(strings.Replace(in, ",", ".", -1), "%f", &out)
	return
}

func main() {
	cgreader.RunAndValidateManualProgram(
		"../../input/defibrillator_1.txt",
		"../../output/defibrillator_1.txt",
		true,
		func(input <-chan string, output chan string) {
			var longitude, latitude float64
			var n int

			longitude, latitude = ParseFloat(<-input), ParseFloat(<-input)
			fmt.Sscanf(<-input, "%d", &n)

			var name string

			for distance, i := math.MaxFloat64, 0; i < n; i++ {
				info := strings.Split(<-input, ";")

				lo, la := ParseFloat(info[4]), ParseFloat(info[5])

				x := (lo - longitude) * math.Cos((latitude+la)/2.0)
				y := la - latitude
				d := math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) * 6371.0

				if d < distance {
					distance, name = d, info[1]
				}
			}

			output <- fmt.Sprintf("%s", name)
		})
}
