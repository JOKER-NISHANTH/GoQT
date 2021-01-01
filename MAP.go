

package main

import "fmt"


func main() {

  
  fmt.Println("Go - User Define Data_Types ")
  fmt.Println(" User Define Data_Types 3. MAP")
  fmt.Println(" 3. MAP is unorder")
  
/*
Syntax
map[key - data_type]value - data_type{key:value}
*/
  MAP :=map[string]int{"Cyber":100,"Nisha":90,"Vasee":95,"Sandy":70}
  // fmt.Printf("%v data-type %T \n",MAP,MAP)
  fmt.Println(MAP)

  // Access the particular values in map
  fmt.Println(MAP["Cyber"])
  
  // Upadte
  MAP["Nandy"]=55
  fmt.Println(MAP)
      
  // Delete
  // Syntax (Variable , key_value)
  delete(MAP,"Cyber") 
  fmt.Println(MAP)

}