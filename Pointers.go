package main

import "fmt"

func main() {

  
  fmt.Println("Go - Pointers")

    ptr := new(int)
    *ptr = 90

    // fmt.Printf("%v \n " ptr)
    fmt.Println(ptr) // Display address
    fmt.Println(*ptr) // Display value

    address :=161
    fmt.Println(&address) // Display address
  
    another_address := &address
    fmt.Println(*another_address)
  
  
  
} 

