package main

import "fmt"


//successive generation of a number
func makeEven()func() uint {
	l := uint(0)

	return func() (ret uint){
		ret = l;
		l +=2
		return;
	}

}
//calculating average
func average (xs []float64)float64{
	total := 0.0 // var total float64 = 0
	for _, value := range xs {
		total += value
	}

	return total / float64(len(xs))
}
//multiple values
func f2()(int,int){
	return 3,67;
}
	//variadic function 
func add(args ...int)int{
	totals := 0;
	for _,value:=range args{
		totals += value;
	}
	return totals
}
 //factorial of n
 func factorial(p uint)uint{
	 if p==0{
		 return 1;
	 }
	 return p * factorial(p-1);
 }
func main (){
	x := []float64{45,56,70.7,56,78,90.65}

	fmt.Println(average(x))
//multiple values
	z,y :=f2()

	fmt.Println(z,y)
//calling add function

sums := add(90,50,60,56,67,89);
fmt.Println(sums)

//using slice and variadic function

arr := []int {90,50,60,56,67,89}
fmt.Println(`using slice: `,add(arr...))


//closure

adds:=func(xi,yi int) int {
	return xi+yi
}
fmt.Println(adds(2,4));

//calling succcessive even generation
nextEven := makeEven()
fmt.Println(nextEven()) // 0
fmt.Println(nextEven()) // 2
fmt.Println(nextEven()) // 4
//calling factorial of f(p)

fmt.Println(factorial(10))

f,_:=os.Open(text.txt)
defer f.close()

}