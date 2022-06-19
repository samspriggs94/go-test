package main

import "fmt"

type slowcar struct {
	name string
}

type fastcar struct {
	name string
}

func (c fastcar) drive() {
	fmt.Println(c.name, "is driving quickly")
}

func (c fastcar) get_name() string {
	return c.name
}

func (c slowcar) drive() {
	fmt.Println(c.name, "is driving slowly")
}

func (c slowcar) get_name() string {
	return c.name
}

// func main() {

// 	// Initializing the values
// 	// of the author structure
// 	slow_car := slowcar{
// 		name: "Ferrari",
// 	}

// 	fast_car := fastcar{
// 		name: "sams car",
// 	}

// 	// Calling the method
// 	slow_car.drive()
// 	fast_car.drive()
// 	drive_the_car(slow_car)
// 	drive_the_car(fast_car)
// }
