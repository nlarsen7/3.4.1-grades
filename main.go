package main

import "fmt"

func gradeA () {
fmt.Println("A")
}
func gradeB () {
  fmt.Println("B")
}
func gradeC () {
  fmt.Println("C")
}
func gradeD () {
  fmt.Println("D")
}
func gradeF () {
  fmt.Println("F")
}
func main () {
var num int
fmt.Println("enter a number between 0-100")
fmt.Scanln(&num)
if num>=90{
  gradeA()
}else if num>=80{
  gradeB()
}else if num>=70{
  gradeC()
}else if num>=60{
  gradeD()
}else{
  gradeF()
}
}