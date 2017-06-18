package main

import "fmt"

func main(){
	var i int
	sum := 0
	fmt.Print("Enter the number: ")
	fmt.Scan(&i)

	for j := 0; j <= i ; j++  {
		sum+=j
	}
	fmt.Print(sum)
}
