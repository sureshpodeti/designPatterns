package abstract

import (
	"abstractFactoryPattern/chair"
	"abstractFactoryPattern/sofa"
	"abstractFactoryPattern/table"
)

type IFurnitureFactory interface {
	CreateChair() chair.IChair
	CreateSofa() sofa.ISofa
	CreateTable() table.ITable
}
