package Hospital

type Department interface {
	Process(p *Patient)
	SendTo(d Department)
}
