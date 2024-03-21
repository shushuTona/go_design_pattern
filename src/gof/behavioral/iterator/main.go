package main

import (
	"fmt"
	"go_design_pattern/src/gof/behavioral/iterator/concrete"
)

func main() {
	ca := &concrete.ConcreteAggregate{}

	ca.Append(
		&concrete.TestData{Name: "Hoge", Age: 20},
		&concrete.TestData{Name: "Piyo", Age: 25},
		&concrete.TestData{Name: "Fuga", Age: 28},
	)

	ci := ca.CreateIterator()
	for ci.HasNext() {
		testData, err := ci.Next()
		if err != nil {
			panic(err)
		}

		testData.Name = fmt.Sprintf("%s_updated", testData.Name)
		fmt.Printf("%v\n", testData)
	}

	fmt.Printf("%#v\n", ca.GetData(0))
	fmt.Printf("%#v\n", ca.GetData(1))
	fmt.Printf("%#v\n", ca.GetData(2))
}
