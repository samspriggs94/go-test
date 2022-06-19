// Golang program to access the interface fields
package main

import "fmt"

// Declare courseprice structure
type Courseprice struct {
	price int
}

// Declare contestprice structure
type Couponprice struct {
	price int
}

type Price interface {
	get_price() int
}

// get_price function for Courseprice
func (a Courseprice) get_price() int {

	return a.price

}

// get_price function for Coupon price
func (b Couponprice) get_price() int {

	return b.price

}

// Compare courseprice and Couponprice.
// Price is interface type
func price_compare(csprice, cpprice Price) bool {

	if csprice.get_price() <= cpprice.get_price() {

		return true

	} else {

		return false

	}

}

func main() {

	var courseprice, couponprice int

	// Get the courseprice from user
	courseprice = 12
	couponprice = 6

	// Create structure of courseprice
	course := Courseprice{courseprice}

	// Create structure of Couponprice
	Coupon := Couponprice{couponprice}

	fmt.Print("Is the course is free: ")

	// Call interface function to compare price
	fmt.Print(price_compare(course, Coupon))
}
