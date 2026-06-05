package main

import "fmt"

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

// func init() {
// 	fmt.Println("Executed Init Function..")
// }

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

// func main() {
// 	mail := make(chan bool)
// 	go func() {
// 		mail <- true
// 	}()
// 	isMailSent := <-mail
// 	if isMailSent {
// 		fmt.Println("Mail was Sent!")
// 	} else {
// 		fmt.Println("Mail was not sent retry")
// 	}
// 	fmt.Println(isMailSent)
// }

/*

Channels

A send operation on an unbuffered channel blocks the sending goroutine
(in our case the main / master go routine)
until another goroutine performs a receive on the same channel, and vice-versa.


If the send operation happens inside the main function,
then the main goroutine is the one that blocks.

In Go, the main function itself runs inside its own goroutine
(often called the master or main goroutine).
It obeys the exact same rules as any other goroutine you create with the go keyword.

Here is a look at what happens under the hood
when the main goroutine gets blocked by an unbuffered channel:

Scenario 1: Total Deadlock (The Program Crashes)
If you try to send data to an unbuffered channel in the main function,
but you haven't started another goroutine to read from it,
the main goroutine blocks forever.

Because Go has a built-in deadlock detector,
it notices that no other goroutines are alive to unblock it,
and it crashes the program.


Scenario 2: Successful Handshake (How to fix it)
To prevent the main goroutine from blocking forever,
you must spawn a separate goroutine before you try sending data.
That way, when the main goroutine blocks on the send line,
the background goroutine is alive and available to pull the data out,
completing the handshake.

Unbuffered Channels (The Default)
When you call make(chan int), you get an unbuffered channel.

Behavior: Unbuffered channels have no capacity to hold data.

Synchronization: A send operation on an unbuffered channel
blocks the sending goroutine until another goroutine
performs a receive on the same channel,
and vice-versa. This behaves like a synchronous handshake.

The Golden Rule of Unbuffered Channels
An unbuffered channel can never hold onto data by itself.
Think of it like a baton pass in a relay race.
The runner (the sender) cannot let go of the baton
until the next runner (the receiver) physically puts their hand on it.
If there is no other runner there, the sender is stuck standing still.


Always Remember : Spawn the Goroutine first and then feed the value

*/

// func main() {
// 	gr1 := make(chan string)
// 	go func() {
// 		gr1 <- "This message is for Agent 57"
// 	}()
// 	capVal := <-gr1
// 	fmt.Println(capVal)
// }

// func main() {
// 	// gr1 := make(chan int)
// 	// gr1 <- 101
// 	// val1 := <-gr1
// 	// fmt.Println(val1)

// 	gr1 := make(chan int, 1)
// 	gr1 <- 101
// 	//val1 := <-gr1
// 	//fmt.Println(val1)
// 	fmt.Println(<-gr1)
// 	//fmt.Println(<-gr1)
// }

/*

I will accept value only when receiver is ready.
This is why we get error when
we use buffered channel from the main function

unbuffered channel has no storage space,
that is why sending without having receiver
causes deadlock

however buffered or unbuffered
channel does not hold any values
once it has given the value
to access again declare a variable
and copy that value into that
variable to use it again

*/

// func worker(done chan bool) {
// 	fmt.Println("Working")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")

// 	done <- true
// }

// func main() {
// 	done := make(chan bool)

// 	go worker(done)

// 	<-done
// }

// func worker(done chan bool) {
// 	fmt.Println("Worker Thread working...")
// 	time.Sleep(time.Second * 3)
// 	fmt.Println("Worker Go routine has completed it's work")
// 	done <- true
// }

// func main() {
// 	done := make(chan bool)
// 	go worker(done)
// 	var1 := <-done
// 	fmt.Println(var1)
// }

/*

[Main Thread]                             [Background Goroutine]
     |                                              |
     |--- (go worker) ----------------------------->|
     |                                              |-- Prints "working..."
     |-- Hits `<-done` (BLOCKED)                    |-- Sleeps for 1 second
     |   [Waiting...]                               |   [Sleeping...]
     |   [Waiting...]                               |-- Prints "done"
     |   [Waiting...]                               |-- Sends `true` to channel
     |<-- Unblocks via channel ---------------------|
     |
     X (Program ends safely)


*/

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		ch1 <- "Msg for channel 1"
// 	}()

// 	go func() {
// 		ch2 <- "Msg for channel 2"
// 	}()

// 	select {
// 	case msg1 := <-ch1:
// 		fmt.Println(msg1)
// 	case msg2 := <-ch2:
// 		fmt.Println(msg2)
// 	}
// }

/*
Selects are like Switch but for channels
in multiple cases GO chooses one randomly
GO routine scheduler decides which one to choose

Without select, if you wait on one channel,
your program may get stuck there
even though another channel already has data

With Select, Go waits for both channels.
Whichever becomes ready first, that case runs.
*/

// func main() {
// 	validate := make(chan int)
// 	go func() {
// 		var otp int
// 		fmt.Scan(&otp)
// 		validate <- otp
// 	}()

// 	val := <-validate
// 	if val == 1729 {
// 		fmt.Println("The OTP is valid")
// 	} else {
// 		fmt.Println("Invalid OTP")
// 	}

// }

// func main() {
// 	jobs := make(chan int, 5)
// 	done := make(chan bool)

// 	go func() {
// 		for {
// 			j, more := <-jobs
// 			if more {
// 				fmt.Println("job received ", j)
// 			} else {
// 				fmt.Println("received all jobs")
// 				done <- true
// 				return
// 			}
// 		}
// 	}()
// 	for j := 1; j <= 3; j++ {
// 		jobs <- j
// 		fmt.Println("Job Sent ", j)
// 	}
// 	close(jobs)
// 	fmt.Println("All Jobs Sent")
// 	<-done
// }

/*
Classic Synchronisation Example :
this demonstrates proper synchronisation in sequence
*/
// func drivingCar(driving chan bool) {
// 	fmt.Println("Elder son driving Car")
// 	driving <- true
// }

// func parkedCar(driving chan bool, done chan bool) {
// 	<-driving
// 	fmt.Println("Elder son Parked the Car")
// 	done <- true
// }

// func main() {
// 	driving := make(chan bool)
// 	done := make(chan bool)

// 	go drivingCar(driving)
// 	go parkedCar(driving, done)

// 	<-done
// }

/*
we can use ranges over channels,
also keep receving from the channel
so that the buffer is not full
*/
// func main() {
// 	queues := make(chan string, 3)
// 	queues <- "one"
// 	queues <- "two"
// 	queues <- "three"
// 	close(queues)
// 	for e := range queues {
// 		fmt.Println(e)
// 	}
// }

// func main() {
// 	timer := time.NewTimer(2 * time.Second)

// 	fmt.Println("Waiting...")
// 	/*
// 		This is timer as channel
// 		as the timer completes it's time
// 		GO sends the current time into
// 		that channel
// 	*/
// 	fmt.Println(<-timer.C)

// 	fmt.Println("Timer finished")
// }

// func main() {
// 	ticker := time.NewTicker(time.Second * 2)
// 	for i := 1; i <= 30; i++ {
// 		t := <-ticker.C
// 		fmt.Println("Tick at ", t)
// 	}
// }

// Worker Pool Pattern
// func workerPool(id int, jobs <-chan int, res chan<- int) {
// 	for j := range jobs {
// 		fmt.Println("Worker", id, "Started Job ", j)
// 		time.Sleep(time.Second)
// 		fmt.Println("Worker", id, "Finished Job ", j)
// 		res <- j * 2
// 	}
// }

// func main() {
// 	const numjobs = 5
// 	jobs := make(chan int)
// 	results := make(chan int)

// 	for w := 1; w <= 3; w++ {
// 		go workerPool(w, jobs, results)
// 	}

// 	for i := 1; i <= numjobs; i++ {
// 		jobs <- i
// 	}
// 	close(jobs)
// 	for j := 1; j <= numjobs; j++ {

// 	}

// }

// func main() {
// 	chan1 := make(chan bool)
// 	chan1 <- true
// 	go func() {
// 		receive := <-chan1
// 		fmt.Println(receive)
// 	}()
// 	fmt.Println("Hello from main")
// }

// func main() {
// 	chan1 := make(chan bool)
// 	go func() {
// 		receive := <-chan1
// 		fmt.Println(receive)
// 	}()
// 	chan1 <- true
// 	fmt.Println("Hello from main")
// }

/*
Because in an unbuffered channel,
receiving first does not fail.
It simply waits.

Receiving first is okay → receiver waits.
Sending first is also okay only if receiver is already running somewhere.

But if sender and receiver are in same sequential flow,
the first blocking operation stops everything.
*/

/*
Bad code teaches you more than the working code
below is one such example
we are creating dedicated channels
for both sending and receiving boolean messages
if the go routines spawn in the right sequence
code might work as intended but
this is not garunteed at runtime always
because go routine must and should receive
the message by someone from the channel
if it is sent into the channel this is
the rule of thumb for the go routines
*/
// func main() {
// 	ch1 := make(chan bool)
// 	go func() {
// 		ch1 <- true
// 	}()
// 	go func() {
// 		<-ch1
// 		v1 := <-ch1
// 		fmt.Println(v1)
// 	}()
// 	v2 := <-ch1
// 	fmt.Println(v2)
// }

/*
To avoid go routines deadlock
make sure the sender and receiver
and in different channels
*/
// func main() {
// 	ch := make(chan int)
// 	ch <- 10
// 	value := <-ch
// 	fmt.Println(value)
// }

// func main() {
// 	ch1 := make(chan func())
// 	go func() {
// 		ch1 <- func() {
// 			fmt.Println("Sending function to a channel")
// 			fmt.Println(ch1)
// 			time.Sleep(time.Second * 3)
// 		}
// 	}()
// 	go func() {
// 		f1 := <-ch1
// 		f1()
// 	}()
// 	time.Sleep(time.Second * 2)
// }

// func main() {
// 	ch1 := make(chan bool)
// 	go func() {
// 		ch1 <- false
// 	}()
// 	v1 := <-ch1
// 	if v1 == false {
// 		fmt.Println("Value Changed")
// 	} else {
// 		fmt.Println("Value remains same")
// 	}
// }

// func main() {
// 	chan1 := make(chan string)
// 	// This works fine
// 	go func() {
// 		fmt.Println(<-chan1)
// 		fmt.Println("Hello from separeate GO routine")
// 	}()
// 	chan1 <- "Message 247"

// 	time.Sleep(time.Second * 2)
// }

/*
The two main GO routine points you need to keep in mind is
1. For unbuffered channels, sender and receiver must be in different active goroutines ✅
2. Do not send first and receive later in the same goroutine ✅
Always remember : sending first will block the current GO routine
*/

// func main() {
// 	chan1 := make(chan string)
// 	go func() {
// 		chan1 <- "Message 247"
// 		fmt.Println("Hello from separeate GO routine")
// 	}()
// 	fmt.Println(<-chan1)
// 	// Even this shit works fine
// }

/*
This code below partially solves the sync problem
So the channel synchronized only the handover of value,
not the order of print statements after the handover.

Unbuffered channel synchronizes send and receive only.
After send/receive completes,
both goroutines run independently.
If you need print/order control,
use another channel signal or WaitGroup.

sync works perfectly up until the point where boolean value is trnasfered
After that point,
Go scheduler decides which goroutine gets CPU first.

sender waits until receiver is ready
receiver waits until sender is ready
*/

/*
In this code below,
But because ch1 is unbuffered,
it cannot store true boolean value

So the send operation pauses here
until some other goroutine reaches: <-ch1

flow :

goroutine 1 reaches: ch1 <- true
goroutine 1 tries to send true

but no receiver is ready yet
(receiver has to be immediately available,
because it is unbuffered channel
and there is no buffer storage,
to store the value
into temporary storage)

so goroutine 1 waits at this line

goroutine 2 runs
prints: Two
goroutine 2 reaches: <-ch1

now receiver is ready
true is transferred from goroutine 1 to goroutine 2

goroutine 1 is released
goroutine 1 prints: One

ch1 <- true waits until the receive happens.
Only after receive happens, the send is considered complete.
*/
// func main() {
// 	ch1 := make(chan bool)
// 	go func() {
// 		ch1 <- true
// 		fmt.Println("One")
// 	}()
// 	go func() {
// 		fmt.Println("Two")
// 		<-ch1
// 	}()
// 	time.Sleep(time.Second * 1)
// }

/*
Synchronisation without time.sleep and waitgroups
using unbuffered send and receive channels
*/

/*
What below code does :

main creates ch1 and done channels.
goroutine 1 tries to send true into ch1 and waits because ch1 is unbuffered.
goroutine 2 prints "Two", then receives from ch1.
Once goroutine 2 receives, goroutine 1's send is completed.
Then goroutine 1 prints "One", sends true into done.
main receives from done and exits.
*/
// func main() {
// 	ch1 := make(chan bool)
// 	done := make(chan bool)

// 	go func() {
// 		ch1 <- true
// 		fmt.Println("One")
// 		done <- true
// 	}()

// 	go func() {
// 		fmt.Println("Two")
// 		<-ch1
// 	}()
// 	<-done
// }

/*
go routines in this below code
are in-sync without time.sleep
and without wait group

Control flow :
1. main starts
2. main creates ch1 and done

3. main starts goroutine 1
4. main starts goroutine 2

5. main reaches <-done
   main is blocked here

6. goroutine 1 prints "Two"

7. goroutine 1 reaches:
   ch1 <- true

   Because ch1 is unbuffered, it waits until someone receives.
*/
// func main() {
// 	ch1 := make(chan bool)
// 	done := make(chan bool)

// 	go func() {
// 		fmt.Println("one")
// 		ch1 <- true
// 	}()
// 	go func() {
// 		<-ch1
// 		fmt.Println("two")
// 		done <- true
// 	}()
// 	<-done
// }

// func main() {
// 	ch1 := make(chan bool)
// 	ch2 := make(chan bool)
// 	done := make(chan bool)
// 	go func() {
// 		fmt.Println("One")
// 		ch1 <- true
// 	}()
// 	go func() {
// 		<-ch1
// 		fmt.Println("Two")
// 		ch2 <- true
// 	}()
// 	go func() {
// 		<-ch2
// 		fmt.Println("Three")
// 		done <- true
// 	}()
// 	<-done
// }

// func main() {
// 	ch1 := make(chan bool)
// 	ch2 := make(chan bool)
// 	ch3 := make(chan bool)
// 	ch4 := make(chan bool)
// 	done := make(chan bool)

// 	go func() {
// 		fmt.Println("One")
// 		ch1 <- true
// 	}()
// 	go func() {
// 		<-ch1
// 		fmt.Println("Two")
// 		ch2 <- true
// 	}()
// 	go func() {
// 		<-ch2
// 		fmt.Println("Three")
// 		ch3 <- true
// 	}()
// 	go func() {
// 		<-ch3
// 		fmt.Println("Four")
// 		ch4 <- true
// 	}()
// 	go func() {
// 		<-ch4
// 		fmt.Println("Five")
// 		done <- true
// 	}()
// 	<-done
// }

// func main() {
// 	u1 := make(chan bool)
// 	u2 := make(chan bool)
// 	allUsersDone := make(chan bool)
// 	go func() {
// 		fmt.Println("Ram Entered the ATM")
// 		fmt.Println("Ram Using the ATM")
// 		fmt.Println("Ram Exits the ATM")
// 		u1 <- true
// 	}()
// 	go func() {
// 		<-u1
// 		fmt.Println("Shyam Entered the ATM")
// 		fmt.Println("Shyam Using the ATM")
// 		fmt.Println("Shyam Exits the ATM")
// 		u2 <- true
// 	}()
// 	go func() {
// 		<-u2
// 		fmt.Println("Kishan Entered the ATM")
// 		fmt.Println("Kishan Using the ATM")
// 		fmt.Println("Kishan Exits the ATM")
// 		allUsersDone <- true
// 	}()

/*
If you don't use the sleep timer,
the main go routine exits early
resulting in printing the old value of value1 variable

timer function delays the main so that
it can print the updated value
//try commenting the timer to verify
*/
// func main() {
// 	var value1 int = 17
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			value1 = value1 + 1
// 		}
// 	}()
// 	time.Sleep(time.Second * 1)
// 	fmt.Println(value1)
// }

/*
Typical race condition scenario in GO
9k is a perfect value if you also want
to see right output as equally as wrong output
but it also largely depends on the
machine to machine cores and compute power
but the core idea is to avoid taking
very large and very small values for evident data race
*/

// func main() {
// 	var count int = 0
// 	go func() {
// 		for i := 1; i <= 3000; i++ {
// 			count++
// 		}
// 	}()
// 	go func() {
// 		for i := 1; i <= 3000; i++ {
// 			count++
// 		}
// 	}()
// 	go func() {
// 		for i := 1; i <= 3000; i++ {
// 			count++
// 		}
// 	}()
// 	time.Sleep(time.Second * 1)
// 	fmt.Print(count)
// }

/*
This data race problem
is sovlved by wait groups,
and Atomic Counters
*/
// Garunteed expected output

func main() {
	//var count int = 0
	var v atomic.Uint64
	var w sync.WaitGroup
	w.Go(func() {
		for range 3000 {
			v.Add(1)
		}
	})
	w.Go(func() {
		for range 3000 {
			v.Add(1)
		}
	})
	w.Go(func() {
		for range 3000 {
			v.Add(1)
		}
	})
	w.Wait()
	fmt.Print(v.Load())
}

	<-allUsersDone
}
