package main

import "fmt"

type Decorator interface {
	Decorate() string
}

type Banner struct {
	str string
}

func (self *Banner) getString() string {
	return "*" + self.str + "*"
}

type ComposiotionDecorateBanner struct {
	banner *Banner
}

func NewCompositionDecorateBanner(str string) *ComposiotionDecorateBanner {
	return &ComposiotionDecorateBanner{&Banner{str}}
}

func (self *ComposiotionDecorateBanner) Decorate() string {
	return self.banner.getString()
}

func main() {
	var decorator Decorator
	decorator = NewCompositionDecorateBanner("A")

	fmt.Println(decorator.Decorate())
}
