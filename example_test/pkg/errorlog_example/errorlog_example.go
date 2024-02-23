package errorlog_example

import (
	"errors"
	"fmt"

	// "log"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var ErrNotFound = errors.New("employee not found")

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	// log.SetPrefix("apiCallEmployee(): ")
	log.Print("Hey, It's a log!")
	return &employee, nil
}

func getInformation(id int) (*Employee, error) {
	// log.SetPrefix("getInformation(): ")
	log.Print("try getInformation.")
	if id != 1001 {
		return nil, ErrNotFound
	}
	employee, err := apiCallEmployee(1000)
	if err != nil {
		return nil, fmt.Errorf("got an error when getting the employee information: %v", err)
	}
	return employee, nil
}

func ErrorLog_example() {
	// file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// log.SetOutput(file)
	// log.SetPrefix("ErrorLog_example(): ")
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Test! It's a zerolog message!")
	log.Debug().
		Int("EmployeeID", 1001).
		Msg("Getting employee information")
	log.Debug().
		Str("Name", "John").
		Send()

	employee, err := getInformation(1001)
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("NOT FOUND: %v\n", err)
	} else if err != nil {
		fmt.Printf("Error : %v\n", err)
	} else {
		fmt.Print(employee)
	}

	log.Fatal().Str("error", "TEST Fatal").Send()
}
