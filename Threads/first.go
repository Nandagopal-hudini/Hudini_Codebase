// package main

// import (
//     "fmt"
//     "time"
// )

// func f(from string) {
//     for i := 0; i < 3; i++ {
//         fmt.Println(from, ":", i)
//     }
// }

// func main() {

//     f("direct")

//     go f("goroutine")

//     go func(msg string) {
//         fmt.Println(msg)
//     }("going")

//     time.Sleep(time.Second)
//     fmt.Println("done")
// }

// package main

// import "fmt"

// func main() {

//     messages := make(chan string)

//     go func() { messages <- "ping" }()

//     msg := <-messages
//     fmt.Println(msg)
// }

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
