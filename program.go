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
	var number1, number2 = 10, 20
	// status = "grad" // can not change instant
	fmt.Println("---- User interactive -----")
	fmt.Print("input new name : ")
	fmt.Scanf("%s", &name) // like C
	fmt.Printf("Hello %s !\n", name)
	// normal display
	fmt.Printf("Your name is %v | type: %T \n", name, name) // print data by type (like prinf in C)
	fmt.Println("Age =", age)
	fmt.Println("Score =", score)
	fmt.Println("Pass =", isPass)
	fmt.Println(status)
	fmt.Println("-- math -----")
	fmt.Println("add = ", number1+number2)      // can't add mismatch type
	fmt.Println("minus = ", number1-number2)    // can't add mismatch type
	fmt.Println("multiply = ", number1*number2) // can't add mismatch type
	fmt.Println("divide = ", number1/number2)   // can't add mismatch type
	fmt.Println("mod = ", number1%number2)      // can't add mismatch type
	fmt.Println("-- logic operator -----")
	fmt.Println(number1, " equal to ", number2, " is ", number1 == number2) // symbol like python
	fmt.Println("---- condition -----")
	// if else
	if score > 50 {
		fmt.Println("Your grade is great")
	} else {
		fmt.Println("Umm.. Maybe a little practice will do")
	}
	// switch case
	// switch {
	// case score < 50:
	// 	fmt.Println("Umm.. Maybe a little practice will do")
	// case score >= 50:
	// 	fmt.Println("Your grade is great")
	// }

}
