package main

import "fmt"

type drivecar interface {
	drive()
	get_name() string
}

func drive_the_car(input drivecar) {
	fmt.Println(input, " is about to drive...")
	input.drive()
}
