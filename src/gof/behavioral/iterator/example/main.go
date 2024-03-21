package main

import (
	"fmt"
	"go_design_pattern/src/gof/behavioral/iterator/example/aggregates"
	"go_design_pattern/src/gof/behavioral/iterator/example/entities"
	"go_design_pattern/src/gof/behavioral/iterator/interfaces"
)

func CheckResumeID(iterator interfaces.Iterator[*entities.Resume]) {
	for iterator.HasNext() {
		data := iterator.Next()
		if data != nil && data.ResumeID == 100 {
			fmt.Println(*data)
		}
	}
}

func main() {
	rla := aggregates.NewResumeList(
		&entities.Resume{ResumeID: 100, Name: "Hoge"},
		&entities.Resume{ResumeID: 200, Name: "Piyo"},
		&entities.Resume{ResumeID: 300, Name: "Fuga"},
	)
	rli := rla.CreateIterator()
	CheckResumeID(rli)
}
