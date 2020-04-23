package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	stdin, err := os.OpenFile("input.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdin = stdin
	stout, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout = stout

	var list [100]int64
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &list[i])
		fmt.Printf("value:= %d\n", list[i])
	}

}
