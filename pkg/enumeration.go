package enumeration

type Generator interface {
	Generate(Spec)
}

type Spec struct {
}
