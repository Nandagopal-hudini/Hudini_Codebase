
// // Objective:
// // Learn how to send and receive values using channels.
 
// // Task:
 
// // Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
// // The main function should receive the message from the channel and print it.
// // Hints:
// // Use an unbuffered channel for simplicity.
// package main
// import(
// 	"fmt"
// ) 

// func sendMessage(ch chan string ){
// 	ch <-"Hello, World"
// }


//  func main(){
// 	massageChannel := make(chan string)

// 	go sendMessage(massageChannel)


// 	massage := <-massageChannel
// 	fmt.Println(massage)
//  }







// // -------------------------------------------------------------------------------------
 
// // Objective:
// // Learn how to create and use goroutines.
 
// // Task:
 
// // Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// // Ensure that the main function waits for all goroutines to finish before exiting.
// // Hints:
 
// // Use time.Sleep for delays.
// // Use a sync.WaitGroup to wait for all goroutines to complete.

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Function to print numbers from 1 to 5 with a delay of 1 second between each number
// func printNumbers(id int, wg *sync.WaitGroup) {
	 
// 	for i := 1; i <= 5; i++ {
// 		fmt.Printf(" %d: %d\n", id, i)
// 		time.Sleep(1 * time.Second) // Delay of 1 second
// 	}
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup // Create a WaitGroup

// 	// Launch three goroutines
// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1) // Increment the WaitGroup counter
// 		go printNumbers(i, &wg)
// 	}

// 	// Wait for all goroutines to complete
// 	wg.Wait()
// }








// // -------------------------------------------------------------------------------------
 
// // Objective:
// // Understand how to handle multiple senders with a single receiver using channels.
 
// // Task:
 
// // Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// // The main function should receive both messages from the channel and print them.
 
// package main

// import (
// 	"fmt"
// )

// // sendHello sends "Hello" to the channel
// func sendHello(ch chan string) {
// 	ch <- "Hello"
// }

// // sendWorld sends "World" to the channel
// func sendWorld(ch chan string) {
// 	ch <- "World"
// }

// func main() {
// 	// Create an unbuffered channel of strings
// 	messageChannel := make(chan string)

// 	// Start two goroutines to send messages
// 	go sendHello(messageChannel)
// 	go sendWorld(messageChannel)

// 	// Receive and print both messages
// 	for i := 0; i < 2; i++ {
// 		message := <-messageChannel
// 		fmt.Println(message)
// 	}
// }


// // -------------------------------------------------------------------------------------
 
// // Objective:
// // Understand how to use channels for communication between goroutines.
 
// // Task:
 
// // Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// // The main function should receive these numbers from the channel and print them.
// // Hints:
 
// // Use an unbuffered channel for simplicity.
// // Remember to close the channel when done sending values.
 
// package main

// import (
// 	"fmt"
// )

// // generateNumbers sends numbers from 1 to 10 to the provided channel
// func generateNumbers(ch chan int) {
// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// 	close(ch) // Close the channel after sending all numbers
// }

// func main() {
// 	// Create an unbuffered channel of integers
// 	numberChannel := make(chan int)

// 	// Start a goroutine to generate numbers
// 	go generateNumbers(numberChannel)

// 	// Receive and print numbers from the channel
// 	for number := range numberChannel {
// 		fmt.Println(number)
// 	}
// }

// // -------------------------------------------------------------------------------------
 
// // Objective:
// // Learn how to use a buffered channel.
 
// // Task:
 
// // Write a program that creates a buffered channel with a capacity of 3.
// // The main function should send three integers (1, 2, 3) to the buffered channel without blocking.
// // Then, receive and print the integers from the channel.
 

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// Create a buffered channel with a capacity of 3
// 	bufferedChannel := make(chan int, 3)

// 	// Send three integers to the buffered channel without blocking
// 	bufferedChannel <- 1
// 	bufferedChannel <- 2
// 	bufferedChannel <- 3

// 	// Receive and print the integers from the channel
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(<-bufferedChannel)
// 	}
// }

// // -------------------------------------------------------------------------------------
 
// // Objective:
// // Learn how to use the select statement to handle multiple channels.
 
// // Task:
 
// // Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// // The main function should use a select statement to receive messages from both channels and print them.
// // Hints:
 
// // Use two separate channels.
// // Use the select statement inside a loop to receive from the channels.
 
// package main

// import (
// 	"fmt"
// 	"time"
// )

// // sendMessages sends a series of messages to the provided channel
// func sendMessages(ch chan string, message string) {
// 	for i := 0; i < 5; i++ {
// 		ch <- fmt.Sprintf("%s %d", message, i)
		
// 	}
// 	close(ch) // Close the channel when done sending messages
// }

// func main() {
// 	// Create two channels
// 	ch1 := make(chan string) 
// 	ch2 := make(chan string)

// 	// Start two goroutines to send messages
// 	go sendMessages(ch1, "Goroutine 1")
// 	go sendMessages(ch2, "Goroutine 2")

// 	// Use a select statement inside a loop to receive from both channels
// 	for {
// 		select {
// 		case msg, ok := <-ch1:
// 			if ok {
// 				fmt.Println(msg)
// 			} else {
// 				ch1 = nil // Set ch1 to nil to disable this case
// 			}
// 		case msg, ok := <-ch2:
// 			if ok {
// 				fmt.Println(msg)
// 			} else {
// 				ch2 = nil // Set ch2 to nil to disable this case
// 			}
// 		}

// 		// Break the loop if both channels are closed
// 		if ch1 == nil && ch2 == nil {
// 			break
// 		}
// 	}
// }



// // -------------------------------------------------------------------------------------
 
// // Objective:
// // Use sync.WaitGroup to wait for multiple goroutines to complete.
 
// // Task:
 
// // Write a program that launches three goroutines, each printing a different message.
// // Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Function to print a message
// func printMessage(wg *sync.WaitGroup, message string) {
// 	defer wg.Done() // Notify the WaitGroup that this goroutine is done
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(message)
// 		time.Sleep(500 * time.Millisecond) // Simulate work with a delay
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup // Create a WaitGroup

// 	// Launch three goroutines
// 	messages := []string{"Goroutine 1: Hello", "Goroutine 2: World", "Goroutine 3: !"}

// 	for _, msg := range messages {
// 		wg.Add(1) // Increment the WaitGroup counter
// 		go printMessage(&wg, msg)
// 	}

// 	// Wait for all goroutines to complete
// 	wg.Wait()
// 	fmt.Println("All goroutines have completed.")
// }


// -------------------------------------------------------------------------------------
 
// Objective:
// Learn how to use the select statement to handle multiple channels.
 
// Task:
 
// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.

package main

import (
	"fmt"
	
)

// sendMessages sends a message to the provided channel
func sendMessages(ch chan string, message string) {
	ch <- message
}

func main() {
	// Create two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines to send messages
	go sendMessages(ch1, "Message from Goroutine 1")
	go sendMessages(ch2, "Message from Goroutine 2")

	// Use a select statement inside a loop to receive from both channels
	for {
		select {
		case msg, ok := <-ch1:
			if ok {
				fmt.Println(msg)
			} else {
				ch1 = nil 
			}
		case msg, ok := <-ch2:
			if ok {
				fmt.Println(msg)
			} else {
				ch2 = nil 
			}
		}

		// Break the loop if both channels are closed
		if ch1 == nil && ch2 == nil {
			break
		}
	}	
}
