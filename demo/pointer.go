package main

import "fmt"

type Flight struct {
	Number int
}

func main() {
	// pointer
	// var x int = 10

	// print add of x
	// fmt.Printf("x: %p\n", &x)

	// pointer
	// var y *int
	// y = &x

	// fmt.Println("x: ", x)
	// fmt.Println("y: ", y)
	// fmt.Println("value in y: ", *y)

	// *y = 20 // x = 20

	// fmt.Println("x: ", x)
	// fmt.Println("y: ", y)
	// fmt.Println("value in y: ", *y)

	// pointer to struct
	f := Flight{Number: 123}

	var pf *Flight
	pf = &f

	fmt.Printf("f: %#v\n", f)
	fmt.Printf("pf:  %#v\n", pf)
	fmt.Println("----------------------------------------------------------------")
	// fmt.Println("pf.Number: ", (*pf).Number)
	// fmt.Println("pf.Number: ", pf.Number)

	ppf := &Flight{Number: 456}
	fmt.Printf("ppf: %#v\n", ppf)
	Mutate(ppf)
	fmt.Printf("ppf: %#v\n", ppf)

	fmt.Println("----------------------------------------------------------------")
	nomu := Flight{Number: 789}
	fmt.Printf("nomu: %#v\n", nomu)
	NoMutate(nomu)
	fmt.Printf("nomu: %#v\n", nomu)
}

func Mutate(f *Flight) {
	f.Number = 999
	fmt.Println("f.Number: ", f.Number)
}

func NoMutate(f Flight) {
	f.Number = 999
	fmt.Println("NoMutate f.Number: ", f.Number)
}
