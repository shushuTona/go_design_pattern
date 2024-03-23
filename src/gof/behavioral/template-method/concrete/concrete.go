package concrete

import "fmt"

type DisplayString string

func (d DisplayString) Open() {
	fmt.Println("<<")
}

func (d DisplayString) Print() {
	fmt.Println(d)
}

func (d DisplayString) Close() {
	fmt.Println(">>")
}

type DisplayInt int

func (d DisplayInt) Open() {
	fmt.Println("<<")
}

func (d DisplayInt) Print() {
	fmt.Println(d)
}

func (d DisplayInt) Close() {
	fmt.Println(">>")
}
