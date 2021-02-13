package main

import "fmt"

func zero(x int){
	x = 0;
}
//introducing pointers
func zeros(xptr *int){
	*xptr = 0;
}
//creating pointer using new inbuilt function
func one(lptr *int){
	*lptr = 10
}
func main (){
	x := 6;
	zero(x)
	fmt.Println(`without pointers`,x)
//calling function pointers
	zeros(&x)
	fmt.Println(`with pointers`,x)
//call new inbuilt
lo := new(int)
one(lo)
fmt.Println(*lo)
}
