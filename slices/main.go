package main

import (
	"fmt"
	// "sort"
)


func main(){
	// slice1 := []string{"Hello","team","Abebe"}
	// fmt.Println(slice1)

	//creating slice form array

// 	array := [5]int{1,2,3,4,5}

// 	var slice1 = array[:]

// 	slice2 := array[0:2]

// 	slice3 := array[2:5]


// 	fmt.Println(slice1)
// 	fmt.Println(slice2)
// 	fmt.Println(slice3)


// creating slice form already existing slice


	// slice1 := []int{1,2,3,4,5,6,7,8,9,10}



	// slice2 := slice1[1:5]
	// slice3 := slice1[3:7]

	// fmt.Println(slice2)
	// fmt.Println(slice3)







	//creating slices using make function


	// slice1 := make([]int,7,8)

	// slice2 := make([]int,6)//since capacity is optional


	// fmt.Println(slice1)
	// fmt.Println(slice2)


	// iterate over the slice



	// slice := []int{1,2,3,4,5,6,7,8,9}


	// for i:=0;i<9;i++{
	// 	fmt.Println(slice[i])
	// }


	// for index,value := range slice{
	// 	fmt.Printf("%d,%v\n",index,value)

	// }


	// j := 0

	// for range slice{ // for loop asa while loop
	// 	fmt.Println(slice[j])
	// 	j++
	// }


	// appending to slices

	// slice1 := make([]int,2,5)

	// slice1[0] = 10
	// slice1[1] = 20

	// slice1 = append(slice1, 30,40,50,60,70,80)


	// fmt.Printf("%v,%d,%T",slice1,len(slice1),slice1)



	// modify the elements of slices

	// array := [6]int{1,2,3,4,5,6}

	// slice1 := array[:4]
	// fmt.Println("before",slice1)
	// fmt.Println("before",array)


	// slice1[0] = 10


	// fmt.Println("after",array)
	// fmt.Println("after",slice1)

	// sorting of slices

	// slice1 := []string{"abebe","zemene","demere"}

	// slice2 := []int{4,1,40,3,8}


	// sort.Ints(slice2)
	// sort.Strings(slice1)
	// fmt.Println(slice1)
	// fmt.Println(slice2)


	// a := []int{10, 20, 30}

	// fmt.Println("Slice A:", a)

	// fmt.Printf("Slice A Length:%d and Capacity:%d\n", len(a), cap(a))

	// b := make([]int, 5, 10)

	// copy(b, a)// b is the destination and a is the source 

	// fmt.Println("Slice B:", b)

	// fmt.Printf("Slice B Length:%d and Capacity:%d", len(b), cap(b))


	// appending slice into another slice

	// slice1 := []string{"Go","C#","C++","perl"}
	// slice2 :=[]string{"java","pyhton","js"}

	// slice2 = append(slice2,slice1...)
	// fmt.Println(slice2)

	//multidimensional slice

	slice1 := [][]int{
		{1,2,3,4},
		{5,6,7,8},
		{10,100,101},
		{5,5,5,5,5},
	}
	fmt.Println(slice1)

}