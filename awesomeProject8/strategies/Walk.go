package strategies

import (
	"awesomeProject8/models"
	"math"
)

type Walk struct {
	Speed float64
}

func (w *Walk) CalculateTime(source models.Point, destination models.Point) float64 {
	distance := (destination.X-source.X)*(destination.X-source.X) +
		(destination.Y-source.Y)*(destination.Y-source.Y)

	distance = math.Sqrt(distance)
	return distance / w.Speed
}
