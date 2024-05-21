
package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
 
// Main function to display the menu and handle user input
func main() {
 
    reader := bufio.NewReader(os.Stdin)
 
    for {
        fmt.Println("1. Add Task")
        fmt.Println("2. List Tasks")
        fmt.Println("3. Mark Task as Done")
        fmt.Println("4. Delete Task")
        fmt.Println("5. Exit")
        fmt.Print("Enter choice: ")
 
        input, _ := reader.ReadString('\n')
        choice, err := strconv.Atoi(strings.TrimSpace(input))
 
        if err != nil {
            fmt.Println("Invalid choice, please try again.")
            continue
        }
 
    }
}
 