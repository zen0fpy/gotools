package prototype

type Clonable interface {
	Clone() Clonable
}

type Manger struct {
	prototypes map[string]Clonable
}

func NewProtoManager() *Manger {
	return &Manger{
		prototypes: make(map[string]Clonable),
	}
}

func (m *Manger) Get(name string) Clonable {
	return m.prototypes[name]
}

func (m *Manger) Set(name string, c Clonable) {
	m.prototypes[name] = c
}
