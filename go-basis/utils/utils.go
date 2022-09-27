package utils

type DefaultExtsConstraints interface {
	any
}

type DefaultFunc[T DefaultExtsConstraints] func(*T)

type DefaultExtInterface[T DefaultExtsConstraints] interface {
	Apply(o *T)
}

type DefaultExts[T DefaultExtsConstraints] struct {
	F DefaultFunc[T]
}

func (de *DefaultExts[T]) Apply(o *T) {
	de.F(o)
}

func NewDefaultExts[T DefaultExtsConstraints](f DefaultFunc[T]) *DefaultExts[T] {
	return &DefaultExts[T]{F: f}
}
