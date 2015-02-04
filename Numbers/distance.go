/*
* Distance Between Two Cities - Calculates the * distance between
* two cities and allows the user to specify a * unit of distance.
* This program may require finding coordinates * for the cities
* like latitude and longitude.
* fomula (http://www.movable-type.co.uk/scripts/latlong.html)
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	lat float64
	lon float64
}

func distance(p1 Point, p2 Point) (float64) {
	R := 6.371
	l1 := p1.lat * math.Pi / 180
	l2 := p2.lat * math.Pi / 180
	a := (p1.lat - p2.lat) * math.Pi / 180
	b := (p1.lon - p2.lon) * math.Pi / 180
	x := math.Pow(math.Sin(a/2),2) + math.Cos(l1) * math.Cos(l2) * math.Pow(math.Sin(b/2), 2)
	return 2 * R * math.Atan2(math.Sqrt(x), math.Sqrt(1 - x))
} 

func main() {
	p1 := Point{lat:40.7486, lon:-73.9864}
	p2 := Point{40.0, 74.98}
	fmt.Println(distance(p1,p2))
}