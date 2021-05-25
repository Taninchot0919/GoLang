package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Start ...")
	startWithType()
	defaultType()
	loop()
	whileInGo()
	ifElse(1)
	ifElse(9)
	ifElse(11)
	switchCase(1)
	usingArray()
	differentArrSlice()
	maps()
}

func startWithType() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
	fmt.Println("startWithType")
	var booleanType = true
	shortBooleanType := true
	num1 := 1
	num2 := 2
	fmt.Println(num1 + num2)
	fmt.Println(booleanType)
	fmt.Println(shortBooleanType)
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
}

func defaultType() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
	fmt.Println("Default Type")
	var booleanType bool
	var intType int
	var floatType float64
	var stringType string

	fmt.Println(booleanType)
	fmt.Println(intType)
	fmt.Println(floatType)
	fmt.Println(stringType)

	// put value to type
	const intType2 int = 1
	fmt.Println(intType2)
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
}

func loop() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
	fmt.Println("For Loop")
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
}

func whileInGo() {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
	fmt.Println("Whie in GO")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
}

func ifElse(num int) {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
	if num2 := 9; num+1 == 10 {
		fmt.Println(num2)
		fmt.Println("Yeahhhh")
	} else if num+1 < 10 {
		fmt.Println(num2)
		fmt.Println("Try again with upper number")
	} else {
		fmt.Println(num2)
		fmt.Println("Try again with lower number")
	}
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
}

func switchCase(number int) {
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
	switch number {
	case 1:
		fmt.Println("One")
		break
	case 2:
		fmt.Println("One")
		break
	case 3:
		fmt.Println("One")
		break
	case 4:
		fmt.Println("One")
		break
	case 5:
		fmt.Println("One")
		break
	}
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-")
}

func usingArray() {
	arr1 := [3]int{1, 2, 3}
	fmt.Println("Length of array is :", len(arr1))
	for i := 0; i < len(arr1); i++ {
		fmt.Println("Output in array [", i, "] is : ", arr1[i])
	}
}

func usingSlice() {
	slice := make([]string, 3)

	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"

}

func differentArrSlice() {
	arr := [3]int{0, 1, 2}
	slice := make([]int, len(arr))

	copy(slice, arr[:])
	slice = append(slice, 3, 4)
	for i := 0; i < len(slice); i++ {
		fmt.Println("Slice [", i, "] : ", slice[i])
	}

	// เอาค่าของ slice ตั้งแต่ช่องที่ 0 ถึงน้อยกว่า 2
	slice2 := slice[0:2]
	fmt.Println("Value of slice 2 : ", slice2)

	// เอาค่าตั้งแต่ช่องที่ 2 ถึงก่อนหน้า 4
	fmt.Println("Slice with [2:4] : ", slice2[2:4])
}

func maps() {
	m := make(map[string]int)
	m["math"] = 85
	m["sci"] = 92
	m["itfun"] = 100
	fmt.Println("Length of map m is : ", len(m))
	fmt.Println("Value of map m is : ", m)
}
