package main

import "fmt"

func main() {

  
  fmt.Println("Go - Operators")
  fmt.Println("L O")

  num1 := 60
  num2 := 40
  
  fmt.Printf("%v \n",((num1 >= num2 ) && (num1 == num2)))
  fmt.Printf("%v \n",((num1 >= num2 ) || (num1 == num2)))
  fmt.Printf("%v \n",!(num1 >= num2 ))
  
} 

