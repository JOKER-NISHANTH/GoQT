
package main

import "fmt"


type phone struct{
  ram int8
  color string
  price float32
};

type company struct{
  phone
  name string

}




func main() {

  
  fmt.Println("Go - User Define Data_Types ")
  fmt.Println(" User Define Data_Types 1. Structure")

  MI := phone{ram:16,color:"black",price:99990.99}
  fmt.Printf("%v \n",MI)
  fmt.Printf("%v %v %v \n ", MI.ram,MI.color,MI.price )

  
  // Structure in Structure
  Asus :=company{phone:phone{ram:8,color:"blue",price:78650.99},name:"Asus"}
  fmt.Printf("%v \n",Asus)

  fmt.Printf("%v \n",Asus.phone.color)
      
} 
