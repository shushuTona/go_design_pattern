package factory

import (
	"fmt"
	"go_design_pattern/src/gof/creational/factory-method/example/ifactory"
	"go_design_pattern/src/gof/creational/factory-method/example/iproduct"
	"go_design_pattern/src/gof/creational/factory-method/example/product"
)

func NewIDCardFactory() ifactory.IFactory {
	return &Factory{
		IFactoryInner: &IDCardFactory{},
	}
}

type IDCardFactory struct{}

func (f *IDCardFactory) CreateProduct(owner string) iproduct.IProduct {
	return &product.IDCard{
		Owner: owner,
	}
}

func (f *IDCardFactory) RegisterProduct(product iproduct.IProduct) {
	fmt.Printf("%+vを登録しました。\n", product)
}
