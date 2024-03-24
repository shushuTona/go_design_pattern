package interfaces

type ICar interface {
	Drive()
	Stop()
	Charge(int)
	CheckState()
}

type ICarCreator[T ICar] interface {
	ShipCar() T
	ICarCreatorInner[T]
}

type ICarCreatorInner[T ICar] interface {
	MakeCar() T
	Charge(T)
	RegisterList(T)
	ShowSalesList()
}
