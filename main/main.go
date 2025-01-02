package main

import (
	"fmt"

	"github.com/kysuwu/authwitz"
)

func main() {
	checker := authwitz.NewChecker()
	err := checker.Open("arch.zip")
	if err != nil {
		panic(err)
	}
	defer checker.Close()
	succ, err := checker.Try("asd123")
	if err != nil {
		panic(err)
	}
	if succ {
		fmt.Println("password matched")
	} else {
		fmt.Println("password did not match")
	}
}
