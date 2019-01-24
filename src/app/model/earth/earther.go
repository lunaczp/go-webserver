package earth

type Earther struct {
	Name string
}

func New() *Earther {
	earther := &Earther{}
	return earther
}

func (e *Earther) setName(name string) bool {
	e.Name = name
	return true
}


