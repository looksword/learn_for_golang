package concurrent_example

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var quit = make(chan bool)

func Concurrent_example() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string, 10)

	for _, api := range apis {
		// _, err := http.Get(api)
		// if err != nil {
		// 	fmt.Printf("ERROR: %s is down!\n", api)
		// 	continue
		// }
		// fmt.Printf("SUCCESS: %s is up and running!\n", api)
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

	// size := 2
	// ch2 := make(chan string, size)
	// send(ch2, "one")
	// send(ch2, "two")
	// go send(ch2, "three")
	// go send(ch2, "four")
	// fmt.Println("All data sent to the channel ...")
	// for i := 0; i < 4; i++ {
	// 	fmt.Println(<-ch2)
	// }
	// fmt.Println("Done!")

	ch3 := make(chan string, 1)
	send(ch3, "Hello Looksword!")
	read(ch3)

	ch4 := make(chan string)
	ch5 := make(chan string)
	go process(ch4)
	go replicate(ch5)

	for i := 0; i < 2; i++ {
		select {
		case process := <-ch4:
			fmt.Println(process)
		case replicate := <-ch5:
			fmt.Println(replicate)
		}
	}

	start = time.Now()
	size := 15
	ch1 := make(chan string, size)
	for i := 0; i < size; i++ {
		go fib(float64(i), ch1)
	}
	for i := 0; i < size; i++ {
		fmt.Printf(<-ch1)
	}
	elapsed = time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

	start = time.Now()
	command := ""
	data := make(chan int)
	go fib2(data)
	for {
		num := <-data
		fmt.Println(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}
	time.Sleep(1 * time.Second)
	elapsed = time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

// func send(ch chan string, message string) {
// 	ch <- message
// }

func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func fib(number float64, ch chan string) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)
}

func fib2(c chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}
