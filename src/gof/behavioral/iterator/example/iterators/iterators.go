package iterators

// iteratorパターンを使用したベースの構造体
func NewResourceIterator[T any](list []T) *ResourceIterator[T] {
	return &ResourceIterator[T]{
		list:  list,
		index: 0,
	}
}

type ResourceIterator[T any] struct {
	list  []T
	index int
}

func (ri *ResourceIterator[T]) Next() T {
	var data T
	if !ri.HasNext() {
		return data
	}

	data = ri.list[ri.index]
	ri.index++

	return data
}

func (ri *ResourceIterator[T]) HasNext() bool {
	return ri.index < len(ri.list)
}
