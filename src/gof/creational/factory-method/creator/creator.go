package creator

import (
	"fmt"
	"go_design_pattern/src/gof/creational/factory-method/interfaces"
	"go_design_pattern/src/gof/creational/factory-method/product"
)

type CarCreator[T interfaces.ICar] struct {
	interfaces.ICarCreatorInner[T]
}

func (c *CarCreator[T]) ShipCar() T {
	car := c.MakeCar()
	c.Charge(car)
	c.RegisterList(car)
	return car
}

// =================================================
// Prius
// =================================================
func NewPriusCreator() interfaces.ICarCreator[*product.Prius] {
	return &CarCreator[*product.Prius]{
		ICarCreatorInner: &PriusCreator{},
	}
}

type PriusCreator struct {
	SalesList []*product.Prius
}

func (c *PriusCreator) MakeCar() *product.Prius {
	return &product.Prius{}
}

func (c *PriusCreator) Charge(car *product.Prius) {
	car.Charge(100)
}

func (c *PriusCreator) RegisterList(car *product.Prius) {
	c.SalesList = append(c.SalesList, car)
}

func (c *PriusCreator) ShowSalesList() {
	fmt.Printf("%#v\n", c.SalesList)
}

// =================================================
// LEAF
// =================================================
func NewLeafCreator() interfaces.ICarCreator[*product.Leaf] {
	return &CarCreator[*product.Leaf]{
		ICarCreatorInner: &LeafCreator{},
	}
}

type LeafCreator struct {
	SalesList []*product.Leaf
}

func (c *LeafCreator) MakeCar() *product.Leaf {
	return &product.Leaf{}
}

func (c *LeafCreator) Charge(car *product.Leaf) {
	car.Charge(100)
}

func (c *LeafCreator) RegisterList(car *product.Leaf) {
	c.SalesList = append(c.SalesList, car)
}

func (c *LeafCreator) ShowSalesList() {
	fmt.Printf("%#v\n", c.SalesList)
}
