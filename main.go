package main

import (
	"fmt"
	"os"
	"unsafe"
)

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
	closures()
	recursionFunc()
	stackExample()
	txtCreate()
	pointerExample()
	passByDifferent()
	structExample()
	methodStructExample()
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

func intPlusPlus() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//  มันจำตำแหน่งบน Memory ของฟังก์ชันไว้และจะเก็บค่า i ไว้ให้เรา แม้เราจะสั่งมันทำงานจบไปแล้วก็ตาม เมื่อเรียกใช้อีกครั้งทำให้ค่านั้นถูกอัพเดทจากค่าครั้งก่อนที่เก็บไว้
func closures() {
	fmt.Println("-------- Closure --------")
	nextInt := intPlusPlus()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("----------------")
	nextInt = intPlusPlus()
	fmt.Println(nextInt())
	fmt.Println("----------------")
}

// ------ Recursion ------
func fact(n int) int {
	if n == 0 {
		return 1
	}
	// มันจะไปเรียก fact เรื่อยๆจนกว่า n จะเป็น 0
	return n * fact(n-1)
}

func factUsingFor(n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}

func recursionFunc() {
	fmt.Println("------ Recursion ------")
	number := 10
	fmt.Println("Using recursion :", fact(number))
	fmt.Println("Using for loop :", factUsingFor(number))
	fmt.Println("------------")
}

func stackExample() {
	//defer เป็นการซ้อนลงไปใน stack จะทำอันปกติให้เสร็จก่อนแล้วค่อยกลับมาทำใน stack
	defer fmt.Println("Stack #1")
	fmt.Println("------- Stack -------")
	fmt.Println("Hello Stack")
	defer fmt.Println("Stack #2")
	fmt.Println("Hello Stack #2")
	fmt.Println("--------------")
}

func txtCreate() {
	fmt.Println("--------- Create File -----------")
	// เอาไว้ถ้าเกิด error จะมาทำข้างใน
	/* ทฤษฎีอ่านจาก https://medium.com/grean-developers-family/day-5-7-วัน-ฉันจะเขียน-golang-ให้ได้-2830703b1a48 */
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("System Cannot find path specific")
			fmt.Println("--------------------")
		}
	}()

	path := "/Users/arts/test.txt" //กำหนด path ที่เราจะทำการ create
	_, err := os.Stat(path)        //ให้มีค่าว่างกับ err โดยใช้ os.Stat เป็นการเช็คว่า path นั้นมีอยู่จริงมั้ยถ้าไม่มีให้เกิด Error Path Error
	if os.IsNotExist(err) {        //เช็คว่าเกิด error มั้ย
		file, err := os.Create(path) // ให้สร้างไฟลืขึ้นมา
		if err != nil {              //ถ้า error != null ให้ทำข้างใน (ถ้า err เป็น null ก็คือไม่มี error)
			panic(err) //โยน error เพื่อให้โปรแกรมหยุดทำงาน
		}
		defer file.Close() //ถ้าทำงานใน method นี้เสร็จแล้วให้ file.close ด้วย
	}
	fmt.Println("Done Creating file", path)
}

func pointerExample() {
	fmt.Println("--------- Pointer ----------")

	i := 5
	fmt.Println(i)
	// & เอาไว้ดู address(ตำแหน่ง) ในตัวแปรนั้นๆ
	fmt.Println("Memory ", &i)

	i = 10
	fmt.Println(i)
	fmt.Println("Memory ", &i)
	// ที่ทั้ง 2 ตัวเท่ากันก็เพราะว่าเป็นตัวแปรเดียวกันแค่เปลี่ยนค่า

	fmt.Println("------------------------")
	fmt.Println("Value of i :", i)
	fmt.Println("Address of i :", &i)
	fmt.Println("------------------------")

	p := &i                           //เอาตัวแปร p มาเก็บ address ของ i
	fmt.Println("Value of p :", p)    //แสดง ค่าของ p ก็คือ address ของ i
	fmt.Println("Address of p :", &p) // แสดง address ของ p
	fmt.Println("Value of *p :", *p)  //แสดงค่าทืี่ตำแหน่ง p ขี้อยู่ นั่นก็คือค่าของ i
	// * คือค่าของ address นั้น ส่วน & คือแสดง address ของตัวแปร เราสามารถทำ *&i ก็จะได้ค่าของ i ได้
	fmt.Println("------------------------")

	*p = 10 //เปลี่ยนค่าที่ตำแหน่ง p ชี้ไปหา
	fmt.Println("Value of p :", *p)
	fmt.Println("Value of i :", i)
	fmt.Println("Size of p :", unsafe.Sizeof(p), "bytes.")
	fmt.Println("Size of i :", unsafe.Sizeof(i), "bytes.")
	fmt.Println("------------------------")

}

// Pass By Different

func setZeroPassByValue(i int) {
	i = 0
	fmt.Println("Set Zero Pass by value :", i)
}

func setZeroPassByReference(i *int) { //รับค่าของ address เข้ามา
	*i = 0 //นำค่าของ address ของตัวแปรที่รับเข้ามาให้ = 0
	fmt.Println("Set Zero Pass by Reference :", *i)
}

func passByDifferent() {
	i := 5
	fmt.Println("Value of i :", i)
	setZeroPassByValue(i)
	fmt.Println("Value of i after setZeroPassByValue :", i)

	fmt.Println("----------------------------------")

	fmt.Println("Value of i :", i)
	setZeroPassByReference(&i)                                  //เราใส่ address ที่ชี้ไปยังตัวแปร i
	fmt.Println("Value of i after setZeroPassByReference :", i) //จะเห็นได้ว่า i หลักเราก็กลายเป็น 0 แล้ว
	fmt.Println("============================================")
}

// Structure
type developers struct {
	fname   string
	lname   string
	skills  []string
	exp     int
	salary  int
	company string
}

func structExample() {
	fmt.Println()
	fmt.Println("--------- Struct ----------")
	fmt.Println("Before assign :", developers{})

	dev := developers{
		fname:  "Taninchot",
		lname:  "Phuwaloertthiwat",
		skills: []string{"Node js", "Vue", "Css", "Java", "C"},
		exp:    0,
	}
	dev.salary = 15000
	fmt.Println(dev)
	fmt.Println("---------------------------")

	devp := &dev //เราชี้ที่ address เพราะอยากให้ dev เปลี่ยน
	fmt.Println(devp.company)
	devp.company = "ArtZuZu Company"
	fmt.Println("Value of devP :", devp)
	fmt.Println("Value of dev :", dev)
	fmt.Println("===========================")
}

// Method struct Example

type rect struct {
	width, height int
}

// method struct ต้องบอกว่ารับ struct ของตัวไหน จะเป็น pointer หรือ value ก็ได้ แล้วแต่ว่าเราจะออกแบบยังไง
func (r rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func methodStructExample() {
	fmt.Println()
	r := rect{
		width:  10,
		height: 23,
	}
	fmt.Println("Area of r :", r.area())
	fmt.Println("Perim of r :", r.perim())
	fmt.Println("---------------------------")

	rp := &r
	fmt.Println("Area of rp :", rp.area())
	fmt.Println("Perim of rp :", rp.perim())
	fmt.Println("===========================")
	fmt.Println()
}

// Interface Example
