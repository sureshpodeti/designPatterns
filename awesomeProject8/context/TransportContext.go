package context

import (
	"awesomeProject8/models"
	"awesomeProject8/strategies"
)

type TransportContext struct {
	TransportStrategy strategies.ITransportationStrategy
}

func (t *TransportContext) SetContext(transportStrategy strategies.ITransportationStrategy) {
	t.TransportStrategy = transportStrategy
}

func (t *TransportContext) CalculateTime(source models.Point, destination models.Point) float64 {
	return t.TransportStrategy.CalculateTime(source, destination)
}
