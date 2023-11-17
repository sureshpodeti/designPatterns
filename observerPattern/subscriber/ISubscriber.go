package subscriber

type Observer interface {
	Update(int)
	GetID() int
}
