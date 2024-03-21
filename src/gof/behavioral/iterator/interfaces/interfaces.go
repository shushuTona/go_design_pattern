package interfaces

type Aggregate[T any] interface {
	CreateIterator() Iterator[T]
}

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}
