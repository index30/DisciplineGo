package main

import "fmt"

func main(){
	println("First Hello World!!")

	fmt.Print("fmt.Print don't make new line,")
	fmt.Println("but fmt.Println make new line.")

	var input string
	fmt.Print("Please input anything!!!")
	fmt.Scan(&input)
	fmt.Printf("Your input is %s?", input)
}