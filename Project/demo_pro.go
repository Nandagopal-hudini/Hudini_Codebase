[14:42] Sivanandha M S
// package main
// import "fmt"
 
// func main(){
//  fmt.Println("Enter your choice:")
//  var string1 string
//  fmt.Scanln(&string1)
//  switch {
//  case string1==1:
//      fmt.Println()
 
//  }
 
// }
package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
 
// creating addtask struct
type Task struct {
    id          int
    description string
    isDone      bool
}
 
// created a empty slice
var tasks []Task
// create add task function
func AddTask() {
 
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter description:")
    descp, _ := reader.ReadString('\n')
    tasks = append(tasks, Task{description: descp, isDone: false})
 
}
 
// create list task function to display the tasks
func listTask() {
    fmt.Println("_______________")
    if len(tasks) == 0 {
        fmt.Println("No task")
    }
    for i, task := range tasks {
        status := "pending"
        if task.isDone {
            status = "done"
        }
        fmt.Printf("%d %s [%s]\n", i+1, task.description, status)
    }
}
 
// create markasDone function
func markTaskAsDone() {
    listTask()
    if len(tasks) == 0 {
        return
    }
    fmt.Println("_______________")
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the number to mark as done: ")
    mark, _ := reader.ReadString('\n')
    index, err := strconv.Atoi(strings.TrimSpace(mark))
 
    if err != nil || index < 1 || index > len(tasks) {
        fmt.Println("Invalid task number")
        return
    }
    tasks[index-1].isDone = true
    fmt.Println("task marked as done")
}
 
// create function for deleting function
func deleteTask() {
    listTask()
    if len(tasks) == 0 {
        return
    }
    fmt.Println("_______________")
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the number to delete: ")
    mark, _ := reader.ReadString('\n')
    index, err := strconv.Atoi(strings.TrimSpace(mark))
    if err != nil || index < 1 || index > len(tasks) {
        fmt.Println("Invalid task number")
        return
    }
    tasks = append(tasks[:index-1], tasks[index:]...)
    fmt.Println("Task is deleted")
 
}
 
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
        // var cho int
        // fmt.Scanln(&cho)
 
        input, _ := reader.ReadString('\n')
        choice, err := strconv.Atoi(strings.TrimSpace(input))
 
        if err != nil {
            fmt.Println("Invalid choice, please try again.")
            fmt.Println(err.Error())
            continue
        }
        // fmt.Println(choice)
        switch choice {
        case 1:
            AddTask()
        case 2:
            listTask()
        case 3:
            markTaskAsDone()
        case 4:
            deleteTask()
        case 5:
            fmt.Println("Exiting...")
            return
 
        }
    }
}
 