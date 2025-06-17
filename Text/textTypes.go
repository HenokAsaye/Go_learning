package main

import (
	"fmt"
	
)



func main(){
	c  := "Good Morning"
	b := []byte(c)
	fmt.Printf("%v,%T\n",b,b)
	r := '😫'
	fmt.Printf("%v,%T\n",r,r)
}