package array_example

import (
	"encoding/json"
	"fmt"
)

func printcap(upper int) {
	var numbers []int
	for i := 0; i < upper; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func fibonacci(number int) []int {
	if number <= 2 {
		return make([]int, 0) // nil
	}
	fbnqret := make([]int, number)
	fbnqret[0] = 1
	fbnqret[1] = 1
	for i := 2; i < number; i++ {
		fbnqret[i] = fbnqret[i-1] + fbnqret[i-2]
	}
	return fbnqret
}

func calLuoMa(lmstr string) int {
	lmMap := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	var ret int = 0
	val, present := lmMap[lmstr[len(lmstr)-1]]
	if !present {
		panic("Error: The roman numeral has a bad digit")
	}
	ret += val
	var lastChar byte = lmstr[len(lmstr)-1]
	for i := len(lmstr) - 2; i >= 0; i-- {
		_, present := lmMap[lmstr[i]]
		if !present {
			panic("Error: The roman numeral has a bad digit")
		}
		if lmMap[lastChar] > lmMap[lmstr[i]] {
			ret -= lmMap[lmstr[i]]
		} else {
			ret += lmMap[lmstr[i]]
		}
	}
	return ret
}

func Array_example() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("Array_example(): recover", handler)
		}
	}()

	var a [3]int
	a[1] = 10
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])

	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)
	cities2 := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities2)

	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))

	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD)

	var threeD [3][5][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			for k := 0; k < 2; k++ {
				threeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
			}
		}
	}
	fmt.Println("\nAll at once:", threeD)

	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))

	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	quarter2Extended := quarter2[:4]
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))

	printcap(10)

	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2
	if remove < len(letters) {
		fmt.Println("Before", letters, "Remove ", letters[remove])
		letters = append(letters[:remove], letters[remove+1:]...)
		fmt.Println("After", letters)
	}
	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])
	fmt.Println("Slice2", slice2)

	slice1 := letters[0:2]
	slice3 := letters[1:4]
	slice1[1] = "Z"
	fmt.Println("After", letters)
	fmt.Println("Slice2", slice2)
	fmt.Println("Slice3", slice3)

	// studentsAge := map[string]int{
	// 	"john": 32,
	// 	"bob":  31,
	// }
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println(studentsAge)

	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}

	delete(studentsAge, "john")
	fmt.Println(studentsAge)
	delete(studentsAge, "christy")
	fmt.Println(studentsAge)

	studentsAge["john"] = 30
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}
	for _, age := range studentsAge {
		fmt.Printf("Ages %d\n", age)
	}

	employee := Employee{
		Person: Person{
			LastName:  "Doe",
			FirstName: "John",
		},
	}
	fmt.Println(employee)
	employeeCopy := &employee
	employeeCopy.FirstName = "David"
	fmt.Println(employee)

	employees := []Employee{
		Employee{
			Person: Person{
				LastName: "Doe", FirstName: "John",
			},
		},
		Employee{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}
	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)
	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v\n", decoded)

	fbnq_test := fibonacci(7)
	fmt.Println("fbnq", fbnq_test)

	luomastr := "XVII"
	luomanum := calLuoMa(luomastr)
	fmt.Println("罗马字母", luomastr, "罗马数字", luomanum)
}
