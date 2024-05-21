// Q1-https://www.codewars.com/kata/56f69d9f9400f508fb000ba7

//Count the Monkeys!

package kata

func monkeyCount(n int) []int {
   // Your code here, happy coding!
   count:=make([]int,n)
  for i:=0;i<n;i++{
    count[i]=i+1
  }
   return count
}

//-----------------------------------------------

//Q2-https://www.codewars.com/kata/55685cd7ad70877c23000102

//Return Negative

package kata

func MakeNegative(x int) int {
  switch  {
    case x <0: return x
    case x==0: return 0
    case x>0: return -x
  }
  return 0
}

//-------------------------------------------

//Q3-https://www.codewars.com/kata/58ca658cc0d6401f2700045f

//Find Multiples of a Number

package kata

func FindMultiples(integer, limit int) []int {
  // Your code here!
  multiple:=make([]int,limit/integer)
  index:=0
  for i:=integer;i<=limit;i++{
    if i%integer==0 {
      multiple[index]=i
      index++
      
    }
  }
  return multiple
}

//----------------------------------------------------------------------


//Q4-https://www.codewars.com/kata/5513795bd3fafb56c200049e

//Count by X

package kata


func CountBy(x, n int) []int {
  multiple:=make([]int,n)
  for i:=1;i<=n;i++{
    
    multiple[i-1]=i*x
  }
  return multiple
}

//------------------------------------------------------------------------------

//Q5-https://www.codewars.com/kata/57a083a57cb1f31db7000028

//Powers of 2

package kata
import "math"
func PowersOfTwo(n int) []uint64 {
  // your code goes here
  expo:=make([]uint64,n+1)
  for i:=0;i<n+1;i++{
    expo[i]=uint64(math.Pow(2,float64(i)))
  }
  return expo
}

//---------------------------------------------------------------

//Q6-https://www.codewars.com/kata/5bb904724c47249b10000131

//Total amount of points

package kata
import (
  "strings"
  "strconv"
)

func Points(games []string) int {
  // your code here!
  points := 0;
  x,y := 0,0;
  var str []string;
  for i:=0;i<len(games);i++{
    str = strings.Split(games[i],":")
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


//-----------------------------------------------------

//Q7-https://www.codewars.com/kata/55f2b110f61eb01779000053

//Beginner Series #3 Sum of Numbers

package kata


func GetSum(a, b int) int {
  sum:=0
  min:=0
  max:=0
  if a < b{
    min=a
    max=b
  }else{
    min=b
    max=a
  }
  for i:=min;i<=max;i++{
    sum+=i
  }
  return sum
}

//-------------------------------------------------

//Q8-https://www.codewars.com/kata/57cebe1dc6fdc20c57000ac9

//Shortest Word

package kata
import (
  "strings"
)

func FindShort(s string) int {
str := strings.Split(s," ")
  min:=str[0]
  
  for i:=1;i<len(str);i++{
    if len(str[i])<len(min){
      min=str[i]
    }
  }
  return len(min)
  // your code
}

//-------------------------------------------------------------------

//Q9-https://www.codewars.com/kata/5f8341f6d030dc002a69d7e4

//Least Larger

package kata


func LeastLarger(a []int, i int) int {
  index:=1000
  least:=-1
  for j:=0;j<len(a);j++{
    if a[i]<a[j]{
      if a[j]<index{
        index=a[j]
        least=j
      }
    }
  }
  return least
}

//-------------------------------------------------------------------

//Q10-https://www.codewars.com/kata/59342039eb450e39970000a6

//Count Odd Numbers below n

package kata

func OddCount(n int) int{
  
  return n/2
  //your code here
  }

//---------------------------------------------------------------------