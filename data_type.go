package main;
import "fmt";
func main(){
	//read user inputs

	fmt.Print("Enter a number")
	var input float64;
	fmt.Scanf("%f",&input);
	output := input*2

	fmt.Println(output)

	// i := 0;
	// for i<= 100{
	// 	fmt.Println(i)
	// 	i +=1;

	// }

	//beter option
	//fmt.Println("******* If and For loop section ******")
	for s:=0; s<=100;s++{
		if s%2 ==0{
			fmt.Println(s,"even")
		}else{
			fmt.Println(s,"odd")
		}
	}

	//problem solving

	y := 10;
	if y > 10 {
		fmt.Println("Big");
	}else{
		fmt.Println("small")
	}

	//all numbers divisible by 3 1-100;
    fmt.Println ("Numbers divisible by 3 between 1 and 100")
	for z:=1; z <= 100; z++{
		if z % 3 == 0 {
			fmt.Println(z);
		}
	}
	fmt.Println ("Numbers multiples of 3,5 and 3&5")
	for r:=1; r <= 100; r++{
		if r %3 == 0 && r %5 ==0{
			fmt.Println("FizzBuzz");
			//mutiple of 3
		}else if r %3 ==0 && r %5 !=0{ //multiple of 3
			fmt.Println("Fizz");
		
		}else if r % 3 !=0 && r %5 ==0{ // multiple of 5
			fmt.Println("Buzz");
		}
	}

	//arrys,slice,maps

	//arrays
	//defining array
	var xs[5]float64;
	xs[0] = 98;
	xs[1] = 93;
	xs[2] = 77;
	xs[3] = 82;
	xs[4] = 83;

// a[5]float64 
	// var total float64 = 0;
	// for ks:=0; ks<5;ks++{
	// 	total +=xs[ks]
	// }
	// fmt.Println(total/5)


	//replace 5 with len(xs)
	 var total float64 =0;
	 for v :=0;v<len(xs);v++{
		 total += xs[v]
	 }
	 fmt.Println(total/float64(len(xs)))
//replace for loop iterating with rangex

    var totals float64 = 0
	  for _,value := range x { // _ means we dont  need the for loop
		 total += value
	 }
	 fmt.Println(total/float64(len(xs)))

	 ls := make([]float64, 5,8)


	 fmt.Println(`slicing \t`,float64(len(ls)))
}
