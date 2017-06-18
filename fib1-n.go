package main

import "fmt"

func main(){
	i := 0
	j := 1
	fib := 0
	var a int
	fmt.Print("Enter the number: ")
	fmt.Scan(&a)
	fmt.Print(i)
	for k := 0 ; k <= a ; k++{
		fib = i + j
		j = i
		i = fib
		fmt.Print("\n",fib)
	}

}
