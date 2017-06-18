package main

import (
	"fmt"
)

func main(){
	var a int
	fmt.Print("Enter the number between 1 to 7: ")
	fmt.Scanln(&a)

	if a == 1{
		fmt.Print("Sun")
	}else if a == 2 {
		fmt.Print("Mon")
	}else if a == 3{
		fmt.Print("Tue")
	}else if a == 4{
		fmt.Print("Thur")
	}else if a == 5{
		fmt.Print("Wed")
	}else if a == 6 {
		fmt.Print("Fri")
	}else if a == 7{
		fmt.Print("Sat")
	}else {
		fmt.Print("Invalid Entry")
	}


}
