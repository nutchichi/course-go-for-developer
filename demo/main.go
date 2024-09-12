package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func divide(a, b int) int {
	return a / b
}

func getName2(name string, age int) string {
	return fmt.Sprintf("Name: %s, Age: %d", name, age)
}

type myFunc func(int, int) int

type myFunc2 func(string, int) string

func calculate(fn myFunc, a int, b int) int {
	return fn(a, b)
}

// func getUser(name string, age int) (r string, err error) {
// 	r = fmt.Sprintf("Name: %s, Age: %d", name, age)
// 	return
// }

func getUser(fn myFunc2, name string, age int) (r string, err error) {
	return fn(name, age), nil
}

func swap(a, b int) (int, int) {
	return b, a
}

func swap2(a int, b string) (int, string) {
	return a, b
}

func main() {

	a, b := swap(1, 2)
	fmt.Println(a, b)

	age, name := swap2(50, "nutchichi")
	fmt.Println(age, name)

	var result = func(a1, a2 int) int {
		return a1 + a2
	}
	fmt.Println(result(1, 2))

	r1 := calculate(sum, 1, 2)
	fmt.Println(r1)

	r2 := calculate(minus, 30, 23)
	fmt.Println(r2)

	r3 := calculate(divide, 10, 2)
	fmt.Println(r3)

	r4, _ := getUser(getName2, "Nutchichi", 50)
	fmt.Println(r4)

	// result := sum(1, 2)
	// fmt.Println("Hello, Nutchichi")

}
