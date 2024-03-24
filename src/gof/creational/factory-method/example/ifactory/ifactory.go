package ifactory

import "go_design_pattern/src/gof/creational/factory-method/example/iproduct"

type IFactory interface {
	Create(owner string) iproduct.IProduct
	IFactoryInner
}

type IFactoryInner interface {
	CreateProduct(owner string) iproduct.IProduct
	RegisterProduct(product iproduct.IProduct)
}
