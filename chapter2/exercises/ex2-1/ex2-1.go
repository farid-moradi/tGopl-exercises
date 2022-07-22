package main

// watch out the GOPATH. put the tempconv folder in the GOPATH/src directory
import (
	"fmt"

	"tempconv"
)

func main() {
	fmt.Printf("%v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("%v\n", tempconv.CToF(tempconv.AbsoluteZeroC))
	fmt.Printf("%v\n", tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Printf("%v\n", tempconv.FToK(tempconv.Fahrenheit(tempconv.BoilingC)))
}
