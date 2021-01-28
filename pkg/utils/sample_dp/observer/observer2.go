package main

import (
	"fmt"
	"math/rand"
)

type observer interface {
	update() int
}

type numberGenerator struct {
	observers []observer
}

func (self *numberGenerator) AddObserver(observer observer) {
	self.observers = append(self.observers, observer)
}

func (self *numberGenerator) notifyObservers() []int {
	var result []int
	for _, observer := range self.observers {
		result = append(result, observer.update())
	}
	return result
}

type randomNumberGenerator struct {
	*numberGenerator
}

func (self *randomNumberGenerator) getNumber() int {
	return rand.Intn(50)
}

func (self *randomNumberGenerator) Execute() []int {
	return self.notifyObservers()
}

type number interface {
	getNumber() int
}

type DigitObserver struct {
	generator number
}

func (self *DigitObserver) update() int {
	return self.generator.getNumber()
}

func NewRandomNumberGenerator() *randomNumberGenerator {
	return &randomNumberGenerator{&numberGenerator{}}
}

func main() {
	random := NewRandomNumberGenerator()

	o1 := &DigitObserver{random}
	o2 := &DigitObserver{random}

	random.AddObserver(o1)
	random.AddObserver(o2)

	result := random.Execute()
	for _, r := range result {
		fmt.Println(r)
	}
}
