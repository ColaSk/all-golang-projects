package utils

//类型约束
type DefaultExtsConstraints interface {
	any
}

// 定义修改函数类型
type DefaultFunc[T DefaultExtsConstraints] func(*T)

// 接口类型
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
