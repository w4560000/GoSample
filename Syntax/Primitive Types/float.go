package main

import (
	"fmt"
	"math"
	"unsafe"
)

/*  float型別分為兩種

float32
float64

*/
func main() {

	var maxFloat32 float32 = math.MaxFloat32
	var maxFloat64 float64 = math.MaxFloat64

	var test32 float32 = 0.123456789
	var test64 float64 = 0.123456789123456789123456789

	fmt.Printf("float32 長度= %d byte\n最大值(科學記號法)= %e\n最大值(十進制)= %f\n\n", unsafe.Sizeof(maxFloat32), maxFloat32, maxFloat32)
	fmt.Printf("float64 長度= %d byte\n最大值(科學記號法)= %e\n最大值(十進制)= %f\n\n", unsafe.Sizeof(maxFloat64), maxFloat64, maxFloat64)

	fmt.Println(test32)
	fmt.Println(test64)
}
