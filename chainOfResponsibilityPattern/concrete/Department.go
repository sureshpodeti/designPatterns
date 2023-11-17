package concrete

type Department interface {
	Execute(p *Patient)
	SetNext(department Department)
}
