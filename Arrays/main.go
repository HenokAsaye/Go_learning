package main


import "fmt"

func main(){
	//  var names [3] string
	//  names[0] = "Abebe"
	//  names[1] = "lema"
	//  names[2] = "Kebede"
	//  fmt.Println(names[0])
	//  fmt.Println(names[1])
	//  fmt.Println(names[2])


	// names:= [3]string{"Henok","Aklile","Abebe"}

	// fmt.Println(names)
	// fmt.Println(names)


	// for i:=0 ; i<3 ; i++{
	// 	fmt.Println(names[i])
	// }
	// intializing array using elipses
	// a := [...]int{10,20,30}
	// fmt.Println(len(a))

	// a := [5]int{1:10,2:20,4:56}
	// fmt.Println(a)

	// a:= [6]int{1,2,3,4,5,6}


	//  for i:=0;i<6;i++{
	// 	fmt.Println(a[i])
	//  }
	//  for index,value := range a{
	// 	fmt.Println(index,"=>",value)
	//  }


	//  j:=0

	//  for range a{
	// 	fmt.Println(a[j])
	// 	j++
	//  }


	// array2 := a// copy by value
	// array3 :=&a // copy by reference 


	// a[0] = 2


	// fmt.Println(a)
	// fmt.Println(array2)
	// fmt.Println(array3)



	 A := [3][3]string{
		{"Go","perl","python"},
		{"a","b","c"},
		{"h","j","k"},
	 }



	 for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			fmt.Println(A[i][j])
		}
	 }
}