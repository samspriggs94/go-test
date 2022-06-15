package main

import "fmt"

type Car struct{}

type slowcar struct {
	name string
	Car
}

type fastcar struct {
	name string
	Car
}

type carinterface interface {
	drive()
}

func drivecar(car carinterface) {
	car.drive()
}

func (c fastcar) drive() {
	fmt.Println(c.name, "is driving quickly")
}

func (c slowcar) drive() {
	fmt.Println(c.name, "is driving slowly")
}

func main() {

	// Initializing the values
	// of the author structure
	slow_car := slowcar{
		name: "volvo",
	}

	fast_car := fastcar{
		name: "sams car",
	}

	// Calling the method
	slow_car.drive()
	fast_car.drive()
}
