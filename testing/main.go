package main

import "fmt"

func main(){
x := []float64{68,78,90,30,48,}

var total float64 =0;
for _, value:=range x {
	total += value;
}

fmt.Println(total / float64(len(x)))

//methods of creating a slice
xs :=make([]float64,7,20)
fmt.Println(float64(len(xs)))


arr := [5]float64{1,2,3,4,5}
//xp := arr[0:5] //[low:high]

//xp := arr[:4] //same as [0:4]
xp := arr[:] //same as arr[0:len(arr)]

fmt.Println(xp)

//slice
ds := []int{2,5,6,78,89} //array
slices := append(ds,3,8)

fmt.Println(ds,slices);

//copy function
d := []int{30,67,89,89,80,82,60}
slice2 :=make([]int,6)
copy(slice2,d);

fmt.Println(d,slice2)

//slices := appen([]int,3)
}