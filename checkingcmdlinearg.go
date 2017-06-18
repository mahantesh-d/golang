package main

import (
	"os"
	"fmt"
)

func main()  {
	arg := os.Args[1]
	if arg == "hello"{
		fmt.Print(arg," World")
	}else{
		fmt.Print("wrong entry")
	}
}
