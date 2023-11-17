package abstract

import (
	"abstractFactoryPattern/chair"
	"abstractFactoryPattern/sofa"
	"abstractFactoryPattern/table"
)

type VictorianFurnitureFactory struct {
}

func (v *VictorianFurnitureFactory) CreateChair() chair.IChair {
	return &chair.VictorianChair{}
}

func (v *VictorianFurnitureFactory) CreateSofa() sofa.ISofa {
	return &sofa.VictorianSofa{}
}

func (v *VictorianFurnitureFactory) CreateTable() table.ITable {
	return &table.VictorianTable{}
}
