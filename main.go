//Nicholas Larsen
//4/25/2020
//asks user for a number and outputs a grade depending on the number
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
//each function is a grade
func main () {
var num int
fmt.Println("enter a number between 0-100")
fmt.Scanln(&num)
//scans number so we can decide what grade it is
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
//parameters that tell you what grade the number is and uses other functions to output the grade letter
}