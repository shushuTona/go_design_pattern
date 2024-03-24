package factory

import (
	"go_design_pattern/src/gof/creational/factory-method/example/ifactory"
	"go_design_pattern/src/gof/creational/factory-method/example/iproduct"
)

// 最終的にファクトリーとして生成される構造体
// ファクトリーメソッドパターンとしてインスタンス生成構成をCreateメソッドで定義
// （＝具体的なインスタンス生成処理は埋め込んだifactory.IFactoryInnerで定義する）
type Factory struct {
	ifactory.IFactoryInner
}

func (f *Factory) Create(owner string) iproduct.IProduct {
	product := f.CreateProduct(owner)
	f.RegisterProduct(product)
	return product
}
