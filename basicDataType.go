package main

import "fmt" 

func main() {
    var isTrue = true // bool type data 
	var name string = "shohag"
	var roll int = 523443 // other type-> int  int8  int16  int32  int64 uint uint8 uint16 uint32 uint64 uintptr
	var mark float32 = 34.5 // other type->float32 float64

    fmt.Println(isTrue)
	fmt.Println(name)
	fmt.Printf("%d\n", roll)
	fmt.Printf("%f", mark)
}
