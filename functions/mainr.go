package main;

import "fmt";
func main(){
// fmt.Print("Enter a number:");
// var input float64
// fmt.Scanf("%f",&input);
// output := input*2

// fmt.Println(output)
// two ways to create a variable

var x string;
x="Hello world";
 //or
 y:="Hellow 2 world"

 fmt.Println("first line is", x + "Second",y)

 xs:=5;
 xs +=1

 fmt.Println(xs)


 //converting fahrenheit into celsiu=us

  fmt.Print("Enter Temperatures F:  ");
 var input float64;
 fmt.Scanf("%f",&input);
 output:=(input-32)*5/9
 if(output != 20 && output = 36){
	 fmt.Println("Normal temperature")
 }else{
	fmt.Println(" temperature above")

 }
 fmt.Println(output)
}
