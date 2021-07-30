package a

import "fmt"

func f() {
	print("ok") // mark:info
	var tiny int // mark
	fmt.Println(tiny) // mark:info it's show time
	println("debug") // mark:warn this print is not require
	fmt.Println("last line!") // mark typo
}
