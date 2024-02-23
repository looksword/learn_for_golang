package interface_example

import (
	"encoding/json"
	"example_test/pkg/interface_example/geometry"
	"example_test/pkg/interface_example/onlineshop"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func printInformation(s geometry.Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

type customWriter struct{}

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

// type dollars float32

// func (d dollars) String() string {
// 	return fmt.Sprintf("$%.2f", d)
// }

// type database map[string]dollars

// func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	for item, price := range db {
// 		fmt.Fprintf(w, "%s: %s\n", item, price)
// 	}
// }

func Interface_example() {
	t := geometry.Triangle{}
	t.SetSize(3)
	s := geometry.Square{Size: 4}
	fmt.Println("Perimeter (triangle):", t.Perimeter())
	fmt.Println("Perimeter (square):", s.Perimeter())
	t.DoubleSize()
	fmt.Println("(triangle) Size :", t.GetSize(), ",Perimeter :", t.Perimeter())

	us := upperstring("Learning Go!")
	fmt.Println(us)
	fmt.Println(us.Upper())

	ct := geometry.ColoredTriangle{Color: "blue"}
	ct.SetSize(3)
	fmt.Println("Size:", ct.GetSize())
	fmt.Println("Perimeter (colored)", ct.Perimeter())
	fmt.Println("Perimeter (normal)", ct.Triangle.Perimeter())

	var sh geometry.Shape = geometry.Square{Size: 3}
	fmt.Printf("%T\n", sh)
	fmt.Println("Area: ", sh.Area())
	fmt.Println("Perimeter:", sh.Perimeter())

	var s2 geometry.Shape = geometry.Square{Size: 3}
	printInformation(s2)
	c := geometry.Circle{Radius: 6}
	printInformation(c)

	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)

	// resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=3")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// io.Copy(os.Stdout, resp.Body)

	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	writer := customWriter{}
	io.Copy(writer, resp.Body)

	// db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	// log.Fatal(http.ListenAndServe("localhost:8000", db))

	emp := onlineshop.Employee{Account: onlineshop.Account{FirstName: "look", LastName: "sword"}}
	fmt.Println("init employee:", emp)
	emp.ChangeName("looksword", "wu")
	fmt.Println("changed employee:", emp)
	credits, err := emp.AddCredits(123.4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}
	credits, err = emp.RemoveCredits(87.8)
	if err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	} else {
		fmt.Println("New Credits Balance = ", credits)
	}
}
