/*
Inspect types to understand the nature of go elements
*/
package learn

import (
	"fmt"
	"strconv"
	"unsafe"
)

//learn string and []byte

//learn pointer

//learn stuct

type empty interface{}

type example interface {
	notImplemented()
}

func CastType() error {
	one := 1
	var i empty = one
	var float float32
	float = float32(one)
	switch i.(type) {
	default:
		fmt.Printf("Type error!\n")
	case int:
		fmt.Printf("%d\n", i)
	}
	fmt.Printf("%f\n", float)
	// This will panic at run time
	var e example = i.(example)
	fmt.Printf("%d\n", e.(empty).(int))
	return nil
}

func Str2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func ConvertStrToNumber() error {
	var i int
	fmt.Scanf("%d", &i)
	str := strconv.FormatInt(int64(i), 10)
	hex, _ := strconv.ParseInt(str, 16, 64)
	fmt.Printf("%d\n", hex)
	return nil
}

func CovertNumberAndPointer() {
	str := "A Go variable"
	addr := unsafe.Pointer(&str)
	fmt.Printf("The address of str is %d\n", addr)
	str2 := (*string)(addr)
	fmt.Printf("String constructed from pointer: %s \n", *str2)
	address := uintptr(addr)
	address += 4
	// This has undefined behavior!
	str3 := (*string)(unsafe.Pointer(address))
	fmt.Printf("String constructed from pointer: %s \n", *str3)
}
