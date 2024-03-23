package main

import (
	"go_design_pattern/src/gof/behavioral/template-method/abstract"
	"go_design_pattern/src/gof/behavioral/template-method/concrete"
)

func main() {
	ads := abstract.NewAbstractDisplay(concrete.DisplayString("HOGE"))
	ads.Display()

	adi := abstract.NewAbstractDisplay(concrete.DisplayInt(100))
	adi.Display()
}
