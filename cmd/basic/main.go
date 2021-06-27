package main

import (
	"fmt"
	//findprocess "github.com/binRick/go-find-process"
	findprocess "github.com/binRick/go-find-process"
)

func main() {
	fmt.Println("vim-go")
	findprocess.Pids()
}
