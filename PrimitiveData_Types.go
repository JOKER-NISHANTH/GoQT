// Go - Primitive Data_Types

package main

import "fmt"

func main() {

  // fmt.Println("In Go Have Three Types of Variable Declarations ")

  fmt.Println("Go - Primitive Data_Types")

  var(

    INT int = 55 ; // integers here - to + values allowed
    UNINT uint = 99; // unsigned integers here only + values allowed

    /* int type have 4 type it's related to int size bit like int8 is used to 0 - 255 , int16 , int32 , int64*/
    
    
    FLOAT float32 = 55.55; // float
    /* float have two type of bit values 32 and 64 */
    
    
    COMPLEX complex64 = 6+2i
    /* real + imag (i) 
    Complex have two type of bit values 64 128 */


    STRING string = "I Love Myself";
    BOOL bool = true ;   
  )
  fmt.Printf(" %v it  %T data-type \n",INT,INT)
  fmt.Printf(" %v it  %T data-type \n",UNINT,UNINT)
  fmt.Printf(" %v it  %T data-type \n",FLOAT,FLOAT)
  fmt.Printf(" %v it  %T data-type \n",COMPLEX,COMPLEX)
  fmt.Printf(" %v it  %T data-type \n",STRING,STRING)
  fmt.Printf(" %v it  %T data-type \n",BOOL,BOOL)
  
} 


