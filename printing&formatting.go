package main

import "fmt"

func main(){
  var age = 37
  fmt.Print("Hello")
  fmt.Print("World")

  //with new linw
  fmt.Print("Hello \n")
  fmt.Print("World")

  fmt.Println(" I am ", age, "Years Old")

  //Printf(formatted string) 
  var name = "bigcat"
  var car = "benz"
  fmt.Printf("My name is %v and My Car is %v", name, car)
}