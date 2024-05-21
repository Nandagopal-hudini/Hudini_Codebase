package main

import(
	"bufio"
	"fmt"
	"strings"
	"os"
)

//var sentence []string

func input() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the block of text: ")
	text1, _ := reader.ReadString('\n')
	text := strings.TrimSpace(text1)
	words := strings.Split(text, " ")
	//words:=strings.Fields(text)   
	fmt.Println(text)
	//fmt.Println("List of Frequenct\n")

	countStore:=map[string] int{} //initialise map to store count


	for i:=0;i<len(words);i++{
		countStore[words[i]]++   //
	}	
	///for worD,freq:=range countStore{
       // fmt.Printf("'%s' - %d\n",worD,freq)
   // }


		
	}


func main(){
	input()
}


