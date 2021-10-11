package main

import "fmt"

func main() {
	var num1 = 5            // type inferred
	var num2 int            // explicitly typed
	var num3 float32        // floating point variable
	var name string         // string variable
	var raining bool        // boolean variable
	var rates float32 = 4.5 // declared as float32 and then initialized

	fmt.Println(num1)    // 5
	fmt.Println(num2)    // 0
	fmt.Println(num3)    // 0
	fmt.Println(name)    // "" (empty string)
	fmt.Println(raining) // false
	fmt.Println(rates)   // 4.5
}
