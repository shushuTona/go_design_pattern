package main

import "fmt"

type IPrint interface {
	PrintWeak()
	PrintStrong()
}

type Banner string

func (b Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b)
}

func (b Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b)
}

type PrintBanner struct {
	Banner
}

func (p *PrintBanner) PrintWeak() {
	p.Banner.ShowWithParen()
}

func (p *PrintBanner) PrintStrong() {
	p.Banner.ShowWithAster()
}

func main() {
	pb := &PrintBanner{"HOGE"}
	pb.PrintWeak()
	pb.PrintStrong()
}
