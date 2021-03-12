// Go does not have an exponentiation operator but has functions in the standard library for three types, float64, complex128, and big.Int. As of Go 1.3, all are documented to return 1.

package main

import (
	"fmt"
	"math"
	"math/big"
	"math/cmplx"
)

func main(){
	fmt.Println("float64: ", math.Pow(0,0))
	var b big.Int
	fmt.Println("big integer:", b.Exp(&b, &b, nil))
	fmt.Println("complex: " , cmplx.Pow(0, 0))
}
