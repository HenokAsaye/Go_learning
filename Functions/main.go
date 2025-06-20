package main

import (
	// "crypto/x509"
	"fmt"
	// "math"
)


func simpleFunction(){
	fmt.Println("Hello World!")
}


func add(x int,y int)int{
	totalsum := x + y
	// fmt.Println(totalsum)
	return totalsum
}


func rectangle(w int, h int)(int,int){
	perimeter := 2 * (w +h)
	area := w * h

	return area,perimeter
}


//function arguments call by value
	// func swap(c int , d int)int{
	// 	var temp int

	// 	temp = c

	// 	c = d
	// 	d = temp
	// 	return temp
	// }
// //function arguments call by reference
// 	func swap2(x *int,y *int)int{
// 		var te int
// 		te = *x

// 		*x = *y
// 		*y = te
// 		return te
// 	}

func main(){
	// simpleFunction()
	// sum:= add(20,30)
	// fmt.Println(sum)


	// square := func(x float64)float64{
	// 	return math.Sqrt(x)
	// }
	// fmt.Println(square(4))
	//  a,b := rectangle(3,4)
	//  fmt.Println(a)
	//  fmt.Println(b)
	// a:=100
	// b:=200
	// swap(a,b)
	// fmt.Println(a,b)
	// swap2(&a,&b)
	// fmt.Println("after",a,b)

	//anonymous function
	fmt.Printf("Area of square,%v/n", func(s int)int{
		return s*s
	}(10))



}
