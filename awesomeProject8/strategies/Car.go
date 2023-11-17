package strategies

import (
	"awesomeProject8/models"
	"math"
)

type Car struct {
	Speed float64
}

func (c *Car) CalculateTime(source models.Point, destination models.Point) float64 {
	distance := (destination.X-source.X)*(destination.X-source.X) +
		(destination.Y-source.Y)*(destination.Y-source.Y)

	distance = math.Sqrt(distance)
	return distance / c.Speed
}
