package main

import "fmt"

type Flight struct {
	flightCode   string
	flightNumber int
	destination  string
}

func (f Flight) getFlightNumber() string {
	fn := fmt.Sprintf("flightnumber: %s %d", f.flightCode, f.flightNumber)
	return fn
}

func main() {

	f := Flight{"ABC", 1234, "Japan"}
	fmt.Printf("%#v\n", f)
	fmt.Println(f.flightCode)

	f1 := Flight{flightNumber: 6666, flightCode: "FFF"}
	fmt.Printf("%#v\n", f1)

	fmt.Println(f.getFlightNumber())
	fmt.Println(f1.getFlightNumber())
}
