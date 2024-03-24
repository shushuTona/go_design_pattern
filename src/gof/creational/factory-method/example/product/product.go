package product

import "fmt"

type IDCard struct {
	Owner string
}

func (p *IDCard) Use() {
	fmt.Printf("%sのカードを使用します。\n", p.Owner)
}
