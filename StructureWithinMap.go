package main

import "fmt"

type area struct{
  height,width int
}

func main() {

  
  fmt.Println("Go - User Define Data_Types ")

  fmt.Println(" Structure within MAP ")

  Area := make(map[string]area) 
  
  Area["Vasee"]=area{height:54,width:89}
  Area["Surya"]=area{height:30,width:50}

  fmt.Println(Area)


  // How to declar

  Dynamic_Declar_Map := make(map[string]int)
  Dynamic_Declar_Map["Dear"]=29
  fmt.Println(Dynamic_Declar_Map)


  // Check the values are present
  // key, check :=MAP["Nisha"]
  // fmt.Println(key,check)

  key, check := MAP["Bye"]
  fmt.Println(key,check)

} 




