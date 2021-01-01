package main

import "fmt"

func main() {

  
  fmt.Println("Go - Operators")
  fmt.Println("B O")

  num1 := 60 // 00111100
  num2 := 13 // 00001101
  
  fmt.Printf("%v \n",(num1 & num2)) // 12 = 00001100
  fmt.Printf("%v \n",(num1 | num2))  // 61 = 00111101
  fmt.Printf("%v \n",(num1 ^ num2 )) // 49 = 00110001
  fmt.Printf("%v \n",(num1 >> num2 ))
  fmt.Printf("%v \n",(num1 << num2 ))
} 

