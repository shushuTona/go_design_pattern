package concrete

import (
	"errors"
)

type TestData struct {
	Name string
	Age  int
}

// Aggregate
type ConcreteAggregate struct {
	testDataList []*TestData
}

func (ca *ConcreteAggregate) CreateIterator() *ConcreteIterator {
	return NewConcreteIterator(ca.testDataList)
}

func (ca *ConcreteAggregate) Append(data ...*TestData) {
	ca.testDataList = append(ca.testDataList, data...)
}

func (ca *ConcreteAggregate) GetData(index int) *TestData {
	if index >= len(ca.testDataList) {
		return nil
	}

	return ca.testDataList[index]
}

// Iterator
func NewConcreteIterator(list []*TestData) *ConcreteIterator {
	return &ConcreteIterator{
		list:  list,
		index: 0,
	}
}

type ConcreteIterator struct {
	list  []*TestData
	index int
}

func (ci *ConcreteIterator) Next() (*TestData, error) {
	if !ci.HasNext() {
		return nil, errors.New("no item error")
	}

	data := ci.list[ci.index]
	ci.index++
	return data, nil
}

func (ci *ConcreteIterator) HasNext() bool {
	return ci.index < len(ci.list)
}
