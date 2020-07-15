package main

import (
	"fmt"
	"math"
	"runtime"
	"unsafe"
)

/*  int型別分為八種:
有負數:
		int8   [(-2^7) ~ (2^7 - 1)]   = [-128 ~ 127]  共 (2^8) 個數字
		int16  [(-2^15) ~ (2^15 - 1)] = [-32768 ~ 32767]  共 (2^16) 個數字
		int32  [(-2^31) ~ (2^31 - 1)] = [-2,147,483,648 ~ 2,147,483,647]  共 (2^32) 個數字
		int64  [(-2^63) ~ (2^63 - 1)] = [-9,223,372,036,854,775,808 ~ 9,223,372,036,854,775,807]  共 (2^64) 個數字

無負數:
		uint8  [0 ~ (2^8 - 1)]  共 (2^8) 個數字
		uint16 [0 ~ (2^16 - 1)] 共 (2^16) 個數字
		uint32 [0 ~ (2^32 - 1)] 共 (2^32) 個數字
		uint64 [0 ~ (2^64 - 1)] 共 (2^64) 個數字
*/
func main() {

	// int 與 uint 會根據當前os的位元數
	// 如64位元os 宣告int的參數則為int64 、 uint亦同
	var intParam int = 1
	var uintParam int = 1

	var maxInt8, minInt8 int8 = math.MaxInt8, math.MinInt8
	var maxInt16, minInt16 int16 = math.MaxInt16, math.MinInt16
	var maxInt32, minInt32 int32 = math.MaxInt32, math.MinInt32
	var maxInt64, minInt64 int64 = math.MaxInt64, math.MinInt64

	var maxUint8, minUint8 uint8 = math.MaxUint8, 0
	var maxUint16, minUint16 uint16 = math.MaxUint16, 0
	var maxUint32, minUint32 uint32 = math.MaxUint32, 0
	var maxUint64, minUint64 uint64 = math.MaxUint64, 0

	fmt.Println("當前電腦位元架構:" + runtime.GOARCH)
	fmt.Printf("intParam 長度= %d byte, type=%T\n", unsafe.Sizeof(intParam), intParam)
	fmt.Printf("int8  長度= %d byte, 最大值=%d, 最小值=%d, type=%T\n", unsafe.Sizeof(maxInt8), maxInt8, minInt8, maxInt8)
	fmt.Printf("int16 長度= %d byte, 最大值=%d, 最小值=%d, type=%T\n", unsafe.Sizeof(maxInt16), maxInt16, minInt16, maxInt16)
	fmt.Printf("int32 長度= %d byte, 最大值=%d, 最小值=%d, type=%T\n", unsafe.Sizeof(maxInt32), maxInt32, minInt32, maxInt32)
	fmt.Printf("int64 長度= %d byte, 最大值=%d, 最小值=%d, type=%T\n", unsafe.Sizeof(maxInt64), maxInt64, minInt64, maxInt64)

	fmt.Print("\n--------------------------\n")

	fmt.Printf("uintParam 長度= %d byte, type=%T\n", unsafe.Sizeof(uintParam), uintParam)
	fmt.Printf("uint8  長度= %d byte, 最大值=%d, 最小值=%d, type=%T\n", unsafe.Sizeof(maxUint8), maxUint8, minUint8, maxUint8)
	fmt.Printf("uint16 長度= %d byte, 最大值=%d, 最小值=%d, type=%T\n", unsafe.Sizeof(maxUint16), maxUint16, minUint16, maxUint16)
	fmt.Printf("uint32 長度= %d byte, 最大值=%d, 最小值=%d, type=%T\n", unsafe.Sizeof(maxUint32), maxUint32, minUint32, maxUint32)
	fmt.Printf("uint64 長度= %d byte, 最大值=%d, 最小值=%d, type=%T\n", unsafe.Sizeof(maxUint64), maxUint64, minUint64, maxUint64)

	// math extension
	fmt.Println(math.Abs(-123))
}
