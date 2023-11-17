package strategies

import (
	"awesomeProject8/models"
	"math"
)

type Bike struct {
	Speed float64
}

func (b *Bike) CalculateTime(source models.Point, destination models.Point) float64 {
	distance := (destination.X-source.X)*(destination.X-source.X) +
		(destination.Y-source.Y)*(destination.Y-source.Y)

	distance = math.Sqrt(distance)
	return distance / b.Speed
}
