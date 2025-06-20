package main

import (
	"fmt"
	// "sort"
)

func main(){
	// var emptyMap = map[int]string{}
	// fmt.Println("EmpltyMap:",emptyMap)

	// var Map = map[int]string{
	// 	1:"Abebe",
	// 	2:"Kebede",
	// 	3:"lema",
	// 	4:"jiji",
	// }
	// fmt.Println("Map with elemets",Map)


	//intiializing of map using make function

	// var Map = make(map[string]float64)


	// Map["abebe"] = 44.6
	// Map["kebede"] = 22.3

	// fmt.Println("Map:",Map)



	//length of a map

	// var Map = map[int]string{1:"abebe",2:"kebede",3:"lema"}

	// fmt.Println(len(Map))


	// accessing elements os a map

	// fmt.Println(Map[1])// it will pring out the value by using the key

	// Map := map[int]string{
	// 	1:"GO",
	// 	2:"Java",
	// 	3:"C++",
	// 	4:"c#",
	// }
	
	// adding an element to map
	// Map[5] = "JS"


	//overide Map element

	// Map[2] = "Python"

	// delete an element form map

	// delete(Map,3)//,using the mapname and the key as an argument

	// fmt.Println(Map)
	
	//iterating over a map

	// for key,value := range Map{
	// 	fmt.Println(key,":",value)
	// }
	// since map is unordered it wil loop in unordered manner


	//sorting of maps
	// unsortedMap :=map[int]string{
	// 	30:"HELLO",
	// 	20:"pyhton",
	// 	100:"java",
	// 	90:"c++",

	// }
	// keys :=make([]int,0,len(unsortedMap))
	// for i:= range unsortedMap{
	// 	keys = append(keys, i)
	// }
	// fmt.Println(keys)

	// sort.Ints(keys)
	//  for _,values := range keys{
	// 	fmt.Println(values,unsortedMap[values])
	//  }
	// how to delete all the elemets of Map

	// Map :=map[int]string{
	// 	30:"HELLO",
	// 	20:"pyhton",
	// 	100:"java",
	// 	90:"c++",
	// }


	// for i:= range Map{
	// 	delete(Map,i)
	// }
	// fmt.Println(Map)

	//merge two maps


	Map1 := map[int]string{
		1:"pyhton",
		2:"java",
		3:"c++",
	}

	Map2 := map[int]string{
		4:"c#",
		5:"Go",
		6:"Js",
	}

	for i,j := range Map1{
		Map2[i] = j
	}
	fmt.Println(Map2)

}