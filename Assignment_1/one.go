
package main

import (
	"fmt"
)
// All main functions are written here.
func main(){
	list1:=[]float64{1,2,3,4,5}
	fmt.Printf("Q1- Average is:%.1f",avg(list1))
	checkAge(25)
	fmt.Printf("\nQ3- Reverse is:")
	reverseString("Nandu")
	fmt.Printf("\nQ4- Largest Number:%d",largestNum([]int{10,20,23,17,1,25,30}));

	fmt.Printf("\nQ5- counter")
	p:=Counter{}
	p=Counter.increment(p)
	p=Counter.increment(p)
	Counter.display(p)
	p=Counter.decrement(p)
	Counter.display(p)
	p=Counter.reset(p)
	Counter.display(p)
}
//.......
// Objective: Write a function that takes an array of numbers and returns the average. Use loops and basic arithmetic.
// Function signature:


func avg(list1 []float64)float64 {        //list1 is an array of type float and return type is also flaot
	sum:=0.0							  //sum is initialised as float 
	for i:=0;i<len(list1);i++{			  //itrate from first element to length of array
		sum+=list1[i]					  
	}
	return sum/float64(len(list1))
}

// 2. Check Age
// Write a function that takes an age and prints whether the person is a minor, a young adult, or an adult.
// Use conditional statements.


func checkAge(age int) {												// Using switch Statement
	switch {
		case age < 18: fmt.Print("\nQ2- Age:  Minor")					//for each case print output
		case age > 24: fmt.Print("\nQ2- Age: Adult")
		case 18 < age && age< 24: fmt.Print("\nQ3- Age: Young Adult")
	}
}

// 3. Reverse String
// Objective: Create a function that reverses a string. 
//This will demonstrate basic string manipulation and for loops.

func reverseString(word string) {
	wordRune:=[]rune(word)
	for i:= len(wordRune)-1;i>=0;i--{     // itration starts from last element and to first one
		fmt.Printf("%c",wordRune[i]);	  // print each word in backwords		
	}
}

// 4. Find Largest Number
// Objective: Write a function that takes an array of numbers and returns the largest number.
// Use loops and conditional statements to solve the problem.

func largestNum(numList [] int) int{		
		rand:=0									//set rand as 0
		for i:=0; i<len(numList);i++{			//start itration from 1nd element till last
			if rand < numList[i]{				//check if declared is smaller than list element
				rand= numList[i]				//true, then change element
			}
		}
		return rand 						
}


// 5. Simple Counter
// Create an object that acts as a counter with methods to add, subtract, and reset the count.
// Demonstrate object methods and the use of this.

type Counter struct{
	count 		int
}

func (p Counter) increment() Counter {
	p.count++
	return p
}

func (p Counter) decrement() Counter{
	p.count--
	return p
}
func (p Counter) display(){
	fmt.Printf("\n%d",p)
}
func (p Counter) reset() Counter{
	p.count=0
	return p
}






[12:21] Erics Antony
package kata

import (

  "strings"

  "strconv"

  "fmt"

)
 
func Points(games []string) int {

  // your code here!

  points := 0;

  x,y := 0,0;

  var str []string;

  for _,xi := range games {

    str = strings.Split(xi,":");

    x,_ = strconv.Atoi(str[0]);

    y,_ = strconv.Atoi(str[1]);

    if x > y {

      points += 3;

    } else if x < y {

      points += 0;

    } else if x == y {

      points += 1;

    }

  }

  return points

}