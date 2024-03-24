package product

import (
	"fmt"
)

type Prius struct {
	isDriving    bool
	gasoleneTank int
}

func (c *Prius) Drive() {
	if c.gasoleneTank == 0 {
		fmt.Printf("ガソリンが必要です。, %#v\n", c)
		c.isDriving = false
		return
	}

	c.isDriving = true
	if c.gasoleneTank < 10 {
		c.gasoleneTank = 0
	} else {
		c.gasoleneTank -= 10
	}
}

func (c *Prius) Stop() {
	if !c.isDriving {
		fmt.Printf("既に停止中です。, %#v\n", c)
		return
	}

	c.isDriving = false
}

func (c *Prius) Charge(gasolene int) {
	c.gasoleneTank += gasolene
}

func (c *Prius) CheckState() {
	fmt.Printf("isDriving : %t , gasoleneTank : %d\n", c.isDriving, c.gasoleneTank)
}

type Leaf struct {
	isDriving bool
	battery   int
}

func (c *Leaf) Drive() {
	if c.battery == 0 {
		fmt.Printf("給電が必要です。, %#v\n", c)
		c.isDriving = false
		return
	}

	c.isDriving = true
	if c.battery < 10 {
		c.battery = 0
	} else {
		c.battery -= 10
	}
}

func (c *Leaf) Stop() {
	if !c.isDriving {
		fmt.Printf("既に停止中です。, %#v\n", c)
	}

	c.isDriving = false
}

func (c *Leaf) Charge(energy int) {
	c.battery += energy
}

func (c *Leaf) CheckState() {
	fmt.Printf("isDriving : %t , battery : %d\n", c.isDriving, c.battery)
}
