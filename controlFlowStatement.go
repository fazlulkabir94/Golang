package main

import "fmt" 

func main() {
    var isAnikLove bool = false
    if isAnikLove {
    	fmt.Println("Yes she love me")
    } else {
    	fmt.Println("She didn't love me")
    	fmt.Println("please wait......")
    	if !isAnikLove {
    		isAnikLove = true
    	} else {
    		fmt.Println("Love each other")
    	}
    }
    if isAnikLove {
    	fmt.Println("go and marrige her")
    } else {
    	fmt.Println("plase try to understand")
    }   
}
