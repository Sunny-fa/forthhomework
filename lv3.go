package main

import (
	"fmt"
	"time"
)

func main() {
	go check()
	time.Sleep(1*time.Second)
	fmt.Println("啊！！！我太难了！大佬，放过我其他真的不会了")
}


func check (){
	for a := 2; a <= 10000; a++ {
		count := 0
		for b := 2; b < a; b++ {
			if a%b == 0 {
				count++

				break
			}
		}
		if count==0 {
			fmt.Println(a,"是素数" )
		}

	}

	}








