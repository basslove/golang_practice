package main

import "fmt"

type creater interface {
	createProduct(owner string) User
	registerProduct(User)
}

type User interface {
	Use() string
}

type Factory struct {
}

func (self *Factory) Create(factory creater, owner string) User {
	user := factory.createProduct(owner)
	factory.registerProduct(user)
	return user
}

type IDCard struct {
	owner string
}

func (self *IDCard) Use() string {
	return self.owner
}

type IDCardFactory struct {
	*Factory
	owners []*string
}

func (self *IDCardFactory) createProduct(owner string) User {
	return &IDCard{owner}
}

func (self *IDCardFactory) registerProduct(product User) {
	owner := product.Use()
	self.owners = append(self.owners, &owner)
}

func main() {
	assert := []string{"A", "B", "C"}
	fmt.Println(assert)

	factory := &IDCardFactory{}
	products := []User{
		factory.Create(factory, "A"),
		factory.Create(factory, "B"),
		factory.Create(factory, "C"),
	}

	fmt.Println(products)
}
