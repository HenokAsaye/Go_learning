package main


import "fmt"


func main(){
	// var a int = 10
	// var b int = 20

	// if(a<b){
	// 	fmt.Println("A is lessThan B")
	// 	if(a + 10 == b){
	// 		fmt.Println("A is 10 lessthan B")
	// 	}else{
	// 		fmt.Println("I don't know let")
	// 	}
	// }else{
	// 	fmt.Println("A is greaterthan B")
	// }

//expression switch
	// var Today =2
	// switch{
	// 	case Today==1:
	// 		fmt.Println("Monday")
	// 	case Today==2:
	// 		fmt.Println("Tuesday")
	// 	case Today==3:
	// 		fmt.Println("Wednesday")
	// 	case Today ==4:
	// 		fmt.Println("Thursday")
	// 	case Today == 5:
	// 		fmt.Println("Friday")
	// 	default:
	// 		fmt.Println("weekend")
	// }

	//type switch

	var value interface{} = "Go Programming language"

	switch a:=value.(type){
		case int16:
			fmt.Println("Type intger",a)
		case float64:
			fmt.Println("floating number",a)
		case string:
			fmt.Println("string",a)
		default:
			fmt.Println("UnKNOWN")

	}
}