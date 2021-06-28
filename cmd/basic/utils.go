package main

import "fmt"

func Fatal(err error) {
	if err == nil {
		return
	}
	Panic(err)
}

func Panic(e error) {
	if e != nil {
		panic(fmt.Sprintf("%s", e))
	}
}
