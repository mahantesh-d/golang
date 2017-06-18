package main

import (
	"fmt"
	"os"
)

func main()  {
	arg1 := os.Args
	arg2 := os.Args[1:]
	arg3 := os.Args[2:]

	fmt.Print(arg1, "\n")
	fmt.Print("Hello ",arg2)
	fmt.Print(arg3)
}
