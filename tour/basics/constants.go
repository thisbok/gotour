package main

import "fmt"

const PI = 3.14

func main(){
	const World = "세계"
	fmt.Println("Hello", World)
	fmt.Println("Happy", PI, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}