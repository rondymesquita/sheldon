package main

import (
	"fmt"
	"time"

	"github.com/alexflint/go-arg"
)

var args struct {
	ID      int `arg:"required"`
	Timeout time.Duration
}
func main(){
	arg.MustParse(&args)


	fmt.Println("Hello World", args)
}
