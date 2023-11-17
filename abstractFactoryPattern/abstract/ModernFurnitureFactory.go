package abstract

import (
	"abstractFactoryPattern/chair"
	"abstractFactoryPattern/sofa"
	"abstractFactoryPattern/table"
)

type ModernFurnitureFactory struct {
}

func (m *ModernFurnitureFactory) CreateChair() chair.IChair {
	return &chair.ModernChair{}

}

func (m *ModernFurnitureFactory) CreateSofa() sofa.ISofa {
	return &sofa.ModernSofa{}
}

func (m *ModernFurnitureFactory) CreateTable() table.ITable {
	return &table.ModernTable{}
}
