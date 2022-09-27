package utils

type DefaultExtsConstraints interface {
	any
}

type DefaultExtInterface[T DefaultExtsConstraints] interface {
	Apply(o *T)
}

type DefaultExts[T DefaultExtsConstraints] struct {
	F func(*T)
}

func (de *DefaultExts[T]) Apply(o *T) {
	de.F(o)
}

func NewDefaultExts[T DefaultExtsConstraints](f func(*T)) *DefaultExts[T] {
	return &DefaultExts[T]{F: f}
}
