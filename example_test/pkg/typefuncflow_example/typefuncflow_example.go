package typefuncflow_example

import (
	"example_test/pkg/calculator"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}

func updateName(name *string) {
	*name = "looksword"
}

func givemeanumber() int {
	return -1
}

func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func createnewfile(filename string) {
	newfile, error := os.Create(filename)
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	defer newfile.Close()

	if _, error = io.WriteString(newfile, "Learning Go!"); error != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}

	newfile.Sync()
}

func FizzBuzz() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Printf("FizzBuzz")
		case i%3 == 0:
			fmt.Printf("Fizz")
		case i%5 == 0:
			fmt.Printf("Buzz")
		default:
			fmt.Printf("%d", i)
		}
		if i < 100 {
			fmt.Printf(",")
		} else {
			fmt.Printf("\n")
		}
	}
}

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number > 1 {
		return true
	} else {
		return false
	}
}

func primeless(upper int) {
	fmt.Printf("prime < %d: ", upper)
	for number := 1; number <= upper; number++ {
		if findprimes(number) {
			fmt.Printf("%v ", number)
		}
	}
	fmt.Printf("\n")
}

func check_1() {
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		switch {
		case val < 0:
			fmt.Println("No!!!!!!!!!!!")
			panic("You entered a negative!")
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		default:
			fmt.Println("You entered:", val)
		}
	}
}

func highlow(high int, low int) {
	if high < low {
		fmt.Println("Panic!")
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	fmt.Println("Call: highlow(", high, ",", low, ")")

	// highlow(high, low+1)
}

func TypeFuncFlow_example() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("TypeFuncFlow_example(): recover", handler)
		}
	}()

	var integer32 int = 2147483648
	fmt.Println(integer32)

	// var integer uint = -10
	// fmt.Println(integer)

	fmt.Println(math.MaxFloat32, math.MaxFloat64)

	var firstName string = "W"
	lastName := "123"
	updateName(&lastName)
	fmt.Println(firstName, lastName)

	fullName := "looksword \t(test \"code\")\n"
	fmt.Println(fullName)

	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	sum, mul := calc("1", "2")
	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", mul)

	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)

	x := 27
	if x%2 == 0 {
		fmt.Println(x, "is even")
	}

	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	sec := time.Now().Unix()
	r := rand.New(rand.NewSource(sec))
	i := r.Int31n(10)

	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	default:
		fmt.Print("Others...")
	}

	region, continent := location("Irvine")
	fmt.Printf("god works in %s, %s\n", region, continent)

	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's the weekend, time to rest!")
	}

	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "looksword@online.com"
	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}

	tempf := r.Float64()
	switch {
	case tempf > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}

	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}

	sum100 := 0
	for i := 1; i <= 999; i++ {
		sum100 += i
	}
	fmt.Println("sum of 1..999 is", sum100)

	r_num := r.Int63n(15)
	n_num := 0
	for r_num != 5 {
		r_num = r.Int63n(15)
		n_num++
	}
	fmt.Println(n_num, "times to get", r_num)

	var rr_num int32
	nn_num := 0
	for {
		nn_num++
		if rr_num = r.Int31n(10); rr_num == 5 {
			fmt.Println("finish! spend", nn_num, "times")
			break
		}
	}

	sum100_5 := 0
	for num := 1; num <= 12345; num++ {
		if num%5 == 0 {
			continue
		}
		sum100_5 += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum100_5)

	for i := 1; i <= 3; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}

	createnewfile("learnGo.txt")

	highlow(2, 0)

	FizzBuzz()

	primeless(100)

	check_1()
}
