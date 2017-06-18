package main

import "fmt"

func main(){
	i := 0
	j := 1
	fib := 0
	fmt.Print(i)
	for a := 0; a < 100 ; a++  {
		fib = i + j
		j = i
		i = fib
		fmt.Print("\n",fib)
	}

}
