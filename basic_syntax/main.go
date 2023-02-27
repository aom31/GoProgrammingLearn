package main

import "fmt"

func main() {

	//array
	arrayUse()

	//slice or list
	sliceUse()

	//map or dictionary
	mapUse()

	//loop use only for not have while , do while
	loopForUse()

	//function
	functionUse()

	pointerUse()
}

func pointerUse() {
	var x int = 10
	//use & for access address of variable
	fmt.Printf("show address of x = %v \n", &x)
	//declare pointer of type
	var y *int
	y = &x //so y can point to address x
	fmt.Printf("show address pointer y point to address x = %v \n ", y)
	//if want to show value of pointer use * again because type pointer is *type
	fmt.Printf(" value of pointer y is %v \n", *y)

	//change value pointer
	*y = 20
	fmt.Printf("show value pointer y = %v and show value x = %v", *y, x)

}

func functionUse() {
	// func nameFunction(input) outputForReturn{}

	//and can declare anonymouse function
	x := func(a, b int) int {
		sum := a + b
		return sum
	}

	//call anonymouse function
	addResult := x(10, 20)
	fmt.Printf("show result sum int from anonymouse function = %v \n", addResult)

	fmt.Println(" ### end session ###")
}

func arrayUse() {
	//array
	fmt.Println("====Start with array session====")
	//1. declare
	var x [3]int = [3]int{1, 2, 3}
	var y [3]int
	z := [3]int{5, 6, 7}
	fmt.Printf("show arr x : %#v and arr y : %v and arr z : %#v \n", x, y, z)
	// array 2 digit
	v := [3][2]int{}
	b := [3][2]int{
		{1, 2},
		{2, 3},
		{5, 6},
	}
	fmt.Printf("show arr 2 digit : %v show arr with value : %#v %v \n", v, b, b)

	fmt.Println("#####End session Array####")

}

func sliceUse() {
	fmt.Println("====start with slice/list session====")
	//declare
	x := []int{}
	u := []int{5, 6, 5, 4, 8, 9, 8, 8}
	fmt.Printf("show slice : %v  and %v\n", x, u)
	//when you use slice  You can use Append to add value to slice
	x = append(x, 8, 9)
	fmt.Printf("show x slice value when add value with slice :%v \n", x)
	//want to get length of array or slice use len()
	lengthOfX := len(x)
	fmt.Printf("show length of data slice : %v \n", lengthOfX)

	//when want to cut data from slice
	// variable[indexStart:]  it will select from index start to end of array
	startXToend := u[2:]
	fmt.Printf("show slice index start to end : %v \n", startXToend)
	startNToNindex := u[1:4]
	fmt.Printf("show slice index 1 to index 4 : %v \n", startNToNindex)
	startToNend := u[:5]
	fmt.Printf(" show index start to end at index 5 : %v \n", startToNend)

	fmt.Println("#### end session ####")

}

func mapUse() {
	//declare
	// var nameVar map[key_type]value_type
	contro := map[string]string{}
	contro["th"] = "thailand"
	contro["en"] = "english"

	fmt.Printf("show value from key control[\"th\"] use map : %v \n", contro["th"])
	//when use map can get 2 value   value,boolean = map[key_type]
	valueContro, ok := contro["th"] //you can check value have or not with ok, it return bool
	if ok {
		fmt.Printf(" have value in this key  = %v \n", valueContro)
	} else {
		fmt.Println("no value with key ok = false")
	}

	fmt.Println("### end session ####")

}

func loopForUse() {
	// for varN := initValue; varN condition loop; varN change loop step{}
	valueS := []int{10, 5, 6, 30, 1, 9, 47}
	for i := 0; i < len(valueS); i++ {
		fmt.Printf(" show value slice int with loop valueS[%v] = %v \n", i, valueS[i])
	}

	// for range and it will return 2 value is keyIndex,value := range ...
	for keyIndex, valueIn := range valueS {
		fmt.Printf("show key index of range list = %v and value in list = %v \n", keyIndex, valueIn)
	}

	fmt.Println(" ### end session ###")

}
