package main

import "fmt" 

func main() {
    var x int = 3 // print 3
    y := &x //print 3

    fmt.Println(x)
    fmt.Println(*y)

    *y = 34

    fmt.Println(x) //print 34
}
