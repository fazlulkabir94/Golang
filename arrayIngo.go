package main

import "fmt" 

func main() {
    ArrData := [4] string{"bangladesh", "kabir", "Shohag", "Anika"} //One dimensional array
    
    fmt.Println(len(ArrData))
    fmt.Println()
    fmt.Println(ArrData[0])
    
    data := [3][3] int{          //two dimensional array
    	                {2,3,4},
    	                {3,4,5},
    	                {3,4,6},
                    }
    fmt.Println(data)
    fmt.Println(data[0][0])
    
    //Print two dimensional array use for loop
    for i := 0; i < 3; i++ {
    	for j := 0; j < 3; j++ {
    		fmt.Printf("%d ", data[i][j])
    	}
    	fmt.Println("");
    }
}
