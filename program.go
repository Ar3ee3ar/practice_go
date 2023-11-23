// go will correct syntax when save
package main

import "fmt" //input output lib //need to use lib that import
// use ; to when connect command (it'll enter that command to new line)

func main() {
	name := "Arzeezar" // type-inference : assign data type by type of value
	var age int = 25   // static-inference
	var score float32 = 95.8
	var isPass bool = true
	const status string = "Student"
	// status = "grad" // can not change instant
	fmt.Println("Hello Go Programming")
	fmt.Println("My name is", name)
	fmt.Println("Age =", age)
	fmt.Println("Score =", score)
	fmt.Println("Pass =", isPass)
	fmt.Println(status)

}
