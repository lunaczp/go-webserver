package earther

type Earther struct {
	Name string
}

func New() *Earther {
	earther := &Earther{}
	return earther
}

func (e *Earther) SetName(name string) bool {
	e.Name = name
	return true
}
