package main

import (
	"go_design_pattern/src/gof/creational/factory-method/creator"
)

func main() {
	priusCreator := creator.NewPriusCreator()
	priusCar := priusCreator.ShipCar()
	priusCreator.ShowSalesList()
	priusCar.CheckState()
	priusCar.Drive()
	priusCar.Drive()
	priusCar.Drive()
	priusCar.CheckState()
	priusCar.Charge(10)
	priusCar.CheckState()

	leafCreator := creator.NewLeafCreator()
	leafCar := leafCreator.ShipCar()
	leafCreator.ShowSalesList()
	leafCar.CheckState()
	leafCar.Drive()
	leafCar.Drive()
	leafCar.Drive()
	leafCar.CheckState()
	leafCar.Stop()
	leafCar.CheckState()
}
