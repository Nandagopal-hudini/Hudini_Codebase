//----------------------------------------------------------------
//Q-Sum of a sequence

package kata

func SequenceSum(start, end, step int) int {
  sum:=0
  for i:=start; i<=end; i+=step{
    
    sum=sum+i
  }
  
  return sum
}



//------------------------------------------------------------------
//Q-Sort the Vowels!

package kata

import "strings"

func SortVowels(s string) string {
	str := ""
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			str += "|" + string(s[i]) + "\n"
		default:
			str += string(s[i]) + "|\n"
		}
	}
	return strings.TrimSuffix(str, "\n") // Remove trailing newline
}


//----------------------------------------------------------------------------

//Q- Get the Middle Character

package kata

func GetMiddle(s string) string {
  n:=len(s)/2
  if len(s)%2==0{
    return s[n-1 : n+1]
  }else{
    return string (s[n])
  }
  
}

//-------------------------------------------------------------------------------

//Q Complementary DNA
package kata

func DNAStrand(dna string) string {
	complement := make([]rune, len(dna)) 

	for i, nucleotide := range dna {
		switch nucleotide {
		case 'A':
			complement[i] = 'T'
		case 'T':
			complement[i] = 'A'
		case 'C':
			complement[i] = 'G'
		case 'G':
			complement[i] = 'C'
		}
	}

	return string(complement) 
}

//------------------------------------------------

package kata
import (
  "strings"
  "fmt"
)
func High(s string) string {
  sum := 0;
  maxSum := 0;
  maxWord := "";
  secMaxWord := "";
  secSum := 0;
 
 
  alphaScore := map[string]int{
    "a":1,"b":2, "c":3, "d":4 ,"e":5, "f":6, "g":7, "h":8, "i":9, "j":10,
    "k":11, "l": 12, "m": 13, "n": 14, "o": 15, "p":16, "q": 17, "r": 18, "s": 19,
    "t": 20, "u": 21, "v":22, "w": 23, "x": 24, "y": 25, "z": 26,
  }
  words := strings.Split(s, " ");
  for _, word := range words {
    sum = 0;
    for i:= 0; i < len(word); i++ {
      sum += alphaScore[string(word[i])]
    }
    if sum > maxSum {
      secMaxWord = maxWord
      secSum = maxSum
      maxWord = word;
      maxSum = sum
    }
  }
  fmt.Print("Second:"+secMaxWord)
  fmt.Print("First:"+maxWord)
  if secSum == maxSum {
    return secMaxWord
  }
  return maxWord;
 
}