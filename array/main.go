package main

import "fmt"

func main() {
	fmt.Println("Hi welcome to arrays ")
	var Movielist [4]string

	Movielist[0] = "BATMAN"
	Movielist[1] = "SUPERMAN"
	Movielist[2] = "WONDER WOMAN"
	Movielist[3] = "FLASH"
	fmt.Println("Movielist is", Movielist)
	fmt.Println(len(Movielist))

	var wastelist = [3]string{"SPIDERMAN", "IRONMAN", "THOR"}
	fmt.Println("waste list is", wastelist)

}
