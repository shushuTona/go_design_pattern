package main

import "go_design_pattern/src/gof/creational/factory-method/example/factory"

func main() {
	factory := factory.NewIDCardFactory()
	idCard1 := factory.Create("HOGE")
	idCard1.Use()
}
