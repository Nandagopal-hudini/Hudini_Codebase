




//-----------------------------------------------------------------------------------------
package main

import(
	"fmt"
	"os"
	"time"
)



func main() {
    var waitGroup waitGroup
    
 
    for i := 0; i < 3; i++ {
        waitGroup.add()
        go routine1(&waitGroup)
 
    }
    waitGroup.wait()
 
}
func routine1(wg *waitGroup) {
    defer wg.done()
    for i := 1; i <= 5; i++ {
        fmt.Printf("Routine_1: %d\n", i)
        time.Sleep(time.Second)
    }
}
 
type waitGroup struct {
    count int
}
 
func (wg *waitGroup) add() {
    wg.count++
    fmt.Println(wg.count)
}
func (wg waitGroup) wait() {
    for {
        if wg.count == 0 {
            os.Exit(0)
        }
    }
}
func (wg *waitGroup) done() {
    if wg.count != 0 {
        wg.count--
    }
}