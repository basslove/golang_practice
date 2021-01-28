package main

import "fmt"

var instance *singleton = new(singleton)

type singleton struct{}

func GetInstance() *singleton {
	return instance
}


func main() {
	instance1 := GetInstance()
	instance2 := GetInstance()

	if &instance1 == &instance2 {
		fmt.Println("まっち")
	}else{
		fmt.Println("あんまっち")
	}
}