//Three type of declaration

package main

import "fmt"

func main() {

  fmt.Println("In Go Have Three Types of Variable Declarations ")

  fmt.Println("Static Variable Declarations")
  /*
  Syntax 
  
  var variable_name data_type 
  */

  var STR string = "I am Static Type"
  fmt.Printf(" %v data-type is %T ",STR,STR)

  fmt.Println("\n Dynamic Variable Declarations")

  name := "Cyber-Nishanth"
  fmt.Printf(" %v data-type is %T ",name,name)



 fmt.Println("Group of Variable Declaration")

  var (
    greet string = "Good Morning"
    age int = 21
    BOOL bool = true
  )

fmt.Println(greet, " ",age, " ",BOOL," ")


} 



