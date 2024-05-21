//first program

// All programs are run in package main
package main

//only import modules that are going to use, else it will show error.
import (
	"fmt"       // fmt is the module for printing output
	// "math/rand" //from math module, import rand function to print random numbers
	// "time"      // time module is to display time
)

// func swap(x,y string)(string,string){
// 	return y,x
// }


// func main() {
// 	// fmt.Println("hai its", time.Now(), "now.")
// 	// fmt.Println("selected a random number from 1-10 :", rand.Intn(10))
// 	// fmt.Println(25 + 25)
// 	a,b:=swap("Kill","Joy") //here := declaration+assign combined which is 
// 	fmt.Println(a,b)
// 	var ac = "initial"
//     fmt.Printf("%T\n %s",ac,ac)

// }

//For loop

func main(){
	sum:=0
	for i:=0;i<10;i++{
		sum=sum+i
	}
	fmt.Println(sum)
}





