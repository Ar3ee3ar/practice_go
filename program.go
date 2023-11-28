// go will correct syntax when save
package main

import (
	"fmt"                //input output lib //need to use lib that import
	"gobasic/calculator" // calculator func (package)
)

// use ; to when connect command (it'll enter that command to new line)

func showMessage(name string) (int, string) {
	msg := name + " call -> show msg"
	count := 1
	return count, msg
}

func summation(numbers ...int) int { // not fix number of variable
	total := 0
	for _, value := range numbers { // use _ for complete loop
		total += value
	}
	return total
}

// -------- struct product (OOP)-----------------
type Product struct {
	name     string
	price    float64
	category string
	discount int
}

// --------------- main -----------------------
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
	// default:
	// 	fmt.Println("Your grade is great")
	// }
	fmt.Println("---- array -----")
	// var arr_num [4] int =[4] int {10} // arr_num := [4] int {10,20,30,40} // array with initial size
	arr_num := [...]int{100}
	arr_num[0] = 1
	// fmt.Println(arr_num[2]) // index that not assign value is 0 (have initial size) | can't call (not have size)
	fmt.Println(len(arr_num), arr_num[0])
	fmt.Println("---- slices -----") // like array but can change size (dynamic size)
	sl_name := []string{"a", "b"}
	sl_name = append(sl_name, "c")
	fmt.Println(sl_name)
	fmt.Println("---- maps -----")
	// country := map [string] string {} //var country map [string] string // [key_type] value_type
	// country["th"] = "thailand"
	// country["en"] = "england"
	// or
	country := map[string]string{"th": "thailand", "en": "england"}
	fmt.Println(country["th"])
	fmt.Println("---- for loop -----")
	for i := 0; i <= 5; i++ {
		if i == 1 {
			continue // jump that round
		}
		fmt.Printf("round %d\n", i)
	}
	// condition for loop : break | continue
	for index, value := range country {
		fmt.Printf("%s => %s\n", index, value)
	}
	fmt.Println("------- call func -------")
	fmt.Println(showMessage(name))
	fmt.Printf("sum 3 numbers : %d \n", summation(1, 2, 3))
	fmt.Printf("sum 5 numbers : %d \n", summation(1, 2, 3, 4, 5))
	fmt.Println("------- struct -------")
	product1 := Product{name: "apple", price: 50, category: "fruit", discount: 10}
	fmt.Println(product1)
	fmt.Printf("%s price: %f\n", product1.name, product1.price)
	product1.price -= float64(product1.discount)
	fmt.Printf("after discount => %f\n", product1.price)
	fmt.Println("------- custom package -------")
	fmt.Println(calculator.Add(5, 10))
}
