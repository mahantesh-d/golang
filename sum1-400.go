package main

import "fmt"

func main()  {
	sum := 0
	for i := 1 ; i <= 400 ; i++  {
		sum+=i
	}
	fmt.Print(sum)
}
