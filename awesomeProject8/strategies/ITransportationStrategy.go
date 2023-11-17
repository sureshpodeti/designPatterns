package strategies

import "awesomeProject8/models"

type ITransportationStrategy interface {
	CalculateTime(source models.Point, destination models.Point) float64
}
