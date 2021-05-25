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
	rangeExample()
	cal()
	multipleReturn()
	variadicFunction()
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

	// เอาค่าของ slice ตั้งแต่ช่องที่ 0 ถึง 2 แต่ 2 ไม่เข้าเงื่อนไขเลยไม่ได้ปริ้น
	slice2 := slice[0:2]
	fmt.Println("Value of slice 2 : ", slice2)

	// เอาค่าตั้งแต่ช่องที่ 2 ถึงก่อนหน้า ถึง 4 แต่ 4 ไม่เข้าเงื่อนไขเลยไม่ได้ปริ้น
	fmt.Println("Slice with [2:4] : ", slice2[2:4])
}

func maps() {
	m := make(map[string]int)
	m["math"] = 85
	m["sci"] = 92
	m["itfun"] = 100
	fmt.Println("Length of map m is : ", len(m))
	fmt.Println("Value of map m is : ", m)

	// เพราะว่ามีข้อมูลอยู่แล้วก็เลยใส่ได้เลย
	m2 := map[string]int{"math": 85, "sci": 92, "itFund": 100}
	fmt.Println("Value of m2 is : ", m2)

	delete(m2, "math")
	fmt.Println("Value of m2 after delete is : ", m2)

	value, prs := m2["math"]
	fmt.Println("State of math : ", prs)
	fmt.Println("Value of math : ", value)
}

func rangeExample() {
	arr := [3]int{1, 2, 3}
	slice := []int{4, 5, 6}
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	fmt.Println("Value of array slice and map is : ", arr, slice, m)

	fmt.Println("---- Array ----")
	for i, value := range arr {
		fmt.Println("[", i, "] is ", value)
	}
	fmt.Println("---- Slice ----")
	for i, value := range slice {
		fmt.Println("[", i, "] is ", value)
	}
	fmt.Println("---- Map ----")
	for i, value := range m {
		fmt.Println("[", i, "] is ", value)
	}
}

func plusCal(num1 int, num2 int) int {
	return num1 + num2
}

// บอก int ไว้อันสุดท้ายมันก็รู้ว่าทั้งก้อนเป็น int
func prism(length, height, width int) int {
	return length * height * width
}

func cal() {
	fmt.Println("------------------------------------")
	num1, num2 := 12, 19
	// ประกาศตัวแปรว่าเป็น int
	result := plusCal(num1, num2)
	fmt.Println("Result of", num1, "+", num2, "is :", result)

	// อันนี้เอามาใช้ต่อเลยแค่ = เลย
	result = prism(2, 2, 2)
	fmt.Println("Result of prism cal is :", result)
}

func init() {
	fmt.Println("Init...")
}

// x,y คือ args และ num1 num2 คือ return type
func swapValue(x, y int) (num1, num2 int) {
	num1 = y
	num2 = x
	return
}

// msg string คือ args ส่วน string,string คือ return type
func findFirstAndLastChar(msg string) (string, string) {
	return string(msg[0]), string(msg[len(msg)-1])
}

func multipleReturn() {
	fmt.Println("------------------------------")
	fmt.Println("Go has multiple return value")
	a, b := findFirstAndLastChar("Taninchot Phuwaloertthiwatz")
	fmt.Println(a, b)

	num1, num2 := swapValue(1, 5)
	fmt.Println("num1 :", num1)
	fmt.Println("num2 :", num2)
	fmt.Println("------------------------------")
}

func variadicFunction() {
	nums := []int{1, 2, 3, 4, 5, 6}

	sum(1, 2)
	sum(nums[0:5]...)
	sum(nums...)
}

func sum(numbers ...int) {
	fmt.Print(numbers, "=")
	total := 0
	// ปกติเราจะใส่ i แต่ว่าเพราะเราไม่ได้ใช้ค่าอะไรเลยเกี่ยวกับ i จึงใส่ _ แทนไปได้
	for _, num := range numbers {
		total += num
	}
	fmt.Println(total)
}
