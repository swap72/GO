package main

import (
	"fmt"
)

// func init() {
// 	fmt.Println("This function runs even before the Main function")
// }

// func maincd() {
// 	var name string
// 	fmt.Print("Enter your name:")
// 	fmt.Scan(&name)
// 	fmt.Println("How was your day ? ", name)
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter your string: ")
// 	str, _ := reader.ReadString('\n')
// 	fmt.Println(str)
// }

// func main() {

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter a floating point number: ")
// 	str, _ := reader.ReadString('\n')
// 	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(f)
// 	}

// }

// func main() {
// 	func() {
// 		fmt.Println("Keve a singh? darda kyu a?.. WwwaaaiinKhaalSaaaWwaainFateh!")
// 	}()

// }

// func main() {
// 	wkwf := func() {
// 		fmt.Println("Good Life saddi muteyaree!")
// 	}
// 	wkwf()
// }

// func main() {
// 	var num1 float64 = 4.5210
// 	res := float64(9) + num1
// 	fmt.Println(res)
// }

// exploring pointers

// func main() {
// 	g := 32
// 	age := 45
// 	var ptr *int = &age

// 	if ptr == nil {
// 		fmt.Println("P is nil")
// 	} else {
// 		fmt.Print(g)
// 		fmt.Println("Address : ", ptr)
// 	}
// }

// func add(b int) {
// 	b++
// 	fmt.Println(b)
// 	add(b)
// }

// import "fmt"

// func main() {

// 	name := "Swapnil_Mishra"
// 	const name2 = "Swapnil_Mishra"
// 	var g *string = &name

// 	fmt.Println(name2 == name) // true
// 	fmt.Println(g)             // Prints the memory address of 'name'

// 	i := 1

// 	for i <= 30 {
// 		i = i + 1
// 		fmt.Println(i)
// 	}

// }

// func rec(v int) int {
// 	if v == 0 {
// 		fmt.Println(v)
// 		return 1
// 	}
// 	fmt.Println(v)
// 	return rec(v - 1)
// }

// func mulRet() (int, int) {
// 	return 11, 17
// }

// func init() {
// 	fmt.Println("Let us see, if this run first or not..!")
// }

// func vSum(nums ...int) int {
// 	sum := 0
// 	for _, num := range nums {
// 		sum = sum + num
// 	}
// 	return sum
// }

// func main() {
// 	for {
// 		fmt.Println("Standard infinite for loop in Go language")
// 	}
// }

// func main() {
// 	var age int
// 	fmt.Println("Enter your age: ")
// 	fmt.Scan(&age)
// 	if age >= 13 && age <= 19 {
// 		fmt.Println("Welcome to Teen Age")
// 	} else if age >= 20 {
// 		fmt.Println("You are an Adult")
// 	} else {
// 		fmt.Println("Hello kiddo!")
// 	}
// }

// func main() {
// 	var name = Shiv()()
// 	fmt.Println(name)
// }

// func Shiv() func() string {
// 	return func() string {
// 		return "Narayan"
// 	}
// }

// func parity_of_product(x int, y int) bool {
// 	result := x * y
// 	return result%2 == 0
// }

// func main() {
// 	fmt.Println(time.Now())
// 	switch time.Now().Weekday() {
// 	case time.Saturday, time.Sunday:
// 		fmt.Println("It's Weekend")
// 	default:
// 		fmt.Println("It is workday")
// 	}
// }

// func main() {
// 	var tmp = 17
// 	fmt.Println(passByVal(tmp))
// 	fmt.Println(tmp)
// 	/*
// 		classic pass by value example in Go
// 		pass by value means value of
// 		pass by value means copy of the variable will be passed
// 		but the original variable will not change
// 		that is why the output 18 then 17
// 	*/

// }

// func passByVal(x int) int {
// 	x++
// 	return x
// }

// func numRev(x int) int {
// 	rev := 0
// 	cpx := x
// 	for cpx != 0 {
// 		rev = rev*10 + cpx%10
// 		cpx = cpx / 10
// 	}
// 	return rev
// }

// func main() {
// 	var tmp = 12
// 	fmt.Println(tmp)
// 	ptrManipulation(&tmp)
// 	fmt.Println(tmp)
// 	fmt.Println(tmp)
// }
// // simple pointer manipulation
// func ptrManipulation(ptr *int) {
// 	*ptr = 22
// }

// func main() {
// 	func() {
// 		fmt.Println("Example of an Anonymous function")
// 	}()
// }

// func main() {
// 	CaptialCities := make(map[string]string)
// 	CaptialCities["Karnataka"] = "Bengaluru"
// 	CaptialCities["Rajasthan"] = "Jaipur"
// 	CaptialCities["Kerela"] = "Thiruvananthapuram"
// 	CaptialCities["Andhra Pradesh"] = "Amaravathi"
// 	CaptialCities["Assam"] = "Dispur"
// 	fmt.Println(CaptialCities)
// }

// func main() {
// 	const (
// 		zero = iota
// 		one
// 		two
// 		three
// 		four
// 		five
// 		six
// 		seven
// 	)
// 	fmt.Println(seven + six)
// }

// func main() {
// 	const (
// 		low = iota
// 		medium
// 		average
// 	)
// 	fmt.Println((low + average) / 2)
// }

// type Laptop struct {
// 	ram   int
// 	rom   int
// 	cpu   int
// 	cache int
// 	brand string
// }

// func main() {
// 	L1 := Laptop{
// 		rom:   512,
// 		ram:   16,
// 		cache: 8,
// 		brand: "HP",
// 		cpu:   8,
// 	}
// 	fmt.Println(L1.brand)
// }

// type Shape interface {
// 	Area() float64
// }

// type Circle struct {
// 	radius float64
// }

// func (c Circle) Area() float64 {
// 	return 3.14 * c.radius * c.radius
// }

/*
Interface polymorphs on struct types
given that the struct type implements all the methods of the interface
then the struct type is said to be implementing the interface

Basically Struct Should obey the contract of that interface
https://gobyexample.com/interfaces
*/

// func main() {
// 	circle := Circle{radius: 7}
// 	fmt.Println("Area of this Circle is", circle.Area())
// }

// func main() {
// 	res, err := otpCheck(1717)
// 	if err != nil {
// 		fmt.Println("Wrong Passcode!")
// 	} else {
// 		return res
// 	}

// }

// func otpCheck(otp int) error {
// 	if otp != 1729 {
// 		return fmt.Errorf("Wrong OTP")
// 	}
// 	return nil
// }

// Handling Errors in Go
// func main() {
// 	var otp int
// 	// Calling otpCheck and handling the result
// 	message, err := otpCheck(otp)
// 	if err != nil {
// 		fmt.Println("Error:", err) // Prints the error returned from otpCheck
// 	} else {
// 		fmt.Println(message)
// 	}
// }

// // Function to check OTP
// func otpCheck(otp int) (string, error) {
// 	// Taking user input for OTP
// 	fmt.Scanln(&otp)

// 	// Checking the OTP value
// 	if otp == 1729 {
// 		return "Code Accepted", nil
// 	}

// 	// Returning an error with fmt.Errorf for incorrect OTP
// 	return "Wrong Code", fmt.Errorf("invalid OTP, please try again")
// }

// func main() {
// 	fmt.Println(time.Now().Clock())
// 	/* Prints the time in hh mm ss format */
// }

// func main() {
// 	/*
// 		How Go language is pass by value and not pass by reference
// 	*/
// 	x := 5
// 	y := x
// 	y = 2
// 	fmt.Println(x, y)
// 	/*
// 		y changes because y is working on a copy of x
// 		and is pointing to a different referene in the memory

// 	*/
// }

// func main() {
// 	x := 2
// 	x = increment(x)
// 	fmt.Println(x)
// }

// func increment(x int) int {
// 	return x + 1
// }

/*
in Go statement is not equal to expression
statement performs an action
whereas expression evaluates to a value
*/

// Another example of pass by value

// func main() {
// 	sendsSoFar := 500
// 	const sendsToAdd = 5
// 	incrementSends(sendsSoFar, sendsToAdd)
// 	fmt.Println(sendsSoFar)
// }

// func incrementSends(sendsSoFar int, sendsToAdd int) {
// 	sendsSoFar = sendsSoFar + sendsToAdd
// }

/*
Output should be 505 if it was pass by reference
unlike the actual output
*/

/*
In Go functions can return multiple values
if your function returns multiple values then
you should destructure the return values accordingly
like
func() (int, string) {
	return 17, "Swapnil"
}
	you can also ignore the return value using "_"
	so you only get the value which you are interested in
*/

// func main() {
// 	name, _ := exp()
// 	fmt.Println(name)
// }

// func exp() (string, int) {
// 	return "swapnil", 172
// }

/*

In Go lang we can also name our return values
so we dont have to return them explicitly
In Go, you can give names to return values in the function signature.
meaning :
 	normal return :
	func add(a int, b int) int {
		return a + b
	}

}
	// named return :
	func add(a int, b int) (res int) {
		res = a + b
		return
	}

	named return values hold their initial default values
	eg int : 0
	eg string : ""
	eg bool : false
	eg slice : nil
	eg map : nil
	eg pointer : nil

	also if you use return statement without any return value
	it will automatically return the named return variables with their default values


*/

// func main() {
// 	var alive, name, age = ret()
// 	fmt.Println(alive, name, age)
// }

// func ret() (alive bool, name string, age int) {
// 	return
// }

// type OS interface {
// 	boot() string
// }

// type Windows struct {
// 	hwinfo string
// 	ipk    string
// }

// type Linux struct {
// 	hwinfo  string
// 	verinfo string
// }

// // decoupled service layer
// func boot(o OS) string {
// 	return o.boot()
// }

// // specialized implemented methods for respective types
// func (w Windows) boot() string {
// 	return "Booting Windows | HW: " + w.hwinfo + " | IPK: " + w.ipk
// }

// func (l Linux) boot() string {
// 	return "Booting Linux | HW: " + l.hwinfo + " | Version: " + l.verinfo
// }

// func main() {
// 	// initializing strut fields
// 	win := Windows{hwinfo: "i5 11th Gen", ipk: "9562-2258-5521"}
// 	lin := Linux{hwinfo: "i7 12th Gen", verinfo: "5.17.0-0424-generic"}

// 	fmt.Println(boot(win))
// 	fmt.Println(boot(lin))
// }

// func main() {
// 	a := 172
// 	fmt.Println(memAddr(a))
// }

// Function which returns the memory address of a variable
// func memAddr(a int) *int {
// 	return &a
// }

// GO language allows functions to be called before they are defined
// because GO compiler resolves function declarations at package scope
// meaning order of functions in the source code doesn't matter
// unlike variables

func init() {
	fmt.Println("Executed Init Function..")
}

// type Vehicle interface {
// 	drive()
// }

// type Vehicle interface {
// 	drive()
// }

// type Car struct {
// 	c string
// }

// type Truck struct {
// 	t string
// }

// func (c Car) drive() {
// 	fmt.Println("Driving ", c)
// }

// func (t Truck) drive() {
// 	fmt.Println("Driving ", t)
// }

// // This function wouldn't have been possible if there was no interface
// // This service layer(function) helps loosely decouple the logic
// func Driving(v Vehicle) {
// 	v.drive()
// 	fmt.Println("Driving", v)
// }

// func main() {
// 	c := Car{c: "car"}
// 	t := Truck{t: "truck"}
// 	Driving(c)
// 	Driving(t)
// }

// func otpCheck(otp int) (string, error) {
// 	if otp == 1729 {
// 		return "Code Accepted!", nil
// 	}
// 	return "", errors.New("Invalid OTP, retry later")
// }

/*
this function implements the error handling well in GO.
Now to handle both the return values well,
it is important to caputure both the return values
else both the values will be returned as a tupple
*/

// func main() {
// 	fmt.Println("Enter OTP")
// 	var otp int
// 	fmt.Scan(&otp)
// 	fmt.Println(otpCheck(otp))
// 	// avoid this type of code else it will be returned as tupple
// }

// func otpCheck(otp int) (string, error) {
// 	if otp == 1729 {
// 		return "Code Accepted!", nil
// 	}
// 	return "", errors.New("Invalid OTP, retry later")
// }

// func main() {
// 	fmt.Println("Enter OTP")
// 	var otp int
// 	fmt.Scan(&otp)

// 	// Capture the return values from otpCheck
// 	message, err := otpCheck(otp)

// 	// Check for error and print the appropriate message
// 	if err != nil {
// 		// Print error if OTP is incorrect
// 		fmt.Println("Error:", err)
// 	} else {
// 		// Print success message if OTP is correct
// 		fmt.Println(message)
// 	}
// }

// func main() {
// 	nums := [3]int{11, 17, 67}
// 	fmt.Println(nums[1])
// 	fmt.Println(nums[2])
// 	var o *int = &nums[0]
// 	var p *int = &nums[1]
// 	var q *int = &nums[2]

// 	fmt.Println(o, p, q)
// }

// func cp() {
// 	fmt.Println("Counting Primes..")
// }

// func greeting(func greet(msg string) string) string {
// 	return msg
// }

// func main() {

// 	message, result := bmi(30, func(text string) string {
// 		return text + " Swapnil"
// 	})

// 	fmt.Println(message, result)
// }

// func greet(string) string {
// 	return "Hello Nancy"
// }

// func main() {
// 	var age int = 172
// 	var pointer *int
// 	pointer = &age
// 	fmt.Println(pointer, age)
// }]

// func main() {
// 	go func() {
// 		fmt.Println("This is an anonymous function")
// 	}()
// 	fmt.Println("Hi from Main fn")
// 	time.Sleep(1000)
// }

/*

FLOW :

1. main starts
2. go func() creates a new goroutine
3. main does not wait
4. main prints "Hi from Main fn"
5. goroutine may run after that

// Hi from main fn executed first
// because main does not wait for the goroutine to finish
// it spawns the goroutine and continues executing the next line of code

and using the sleep function we can make sure
that the main function does not exit
before the goroutine has a chance to execute


GO Scheduler is responsible for managing the execution of goroutines
it schedules the goroutines to run on available OS threads
and manages their execution in a way that allows for concurrent execution
Since go scheduler is fast it immidiately
schedules the goroutine for that function and exits the main function

*/
