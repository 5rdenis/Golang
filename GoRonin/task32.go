//Написать собственную функцию Sleep.

package main

import (
    "fmt"
    "time"
)
 

func sleep(x int) {
    <-time.After(time.Second * time.Duration(x))
}  
// Main function
func main() {
  
    // Calling Sleep method
    sleep(5)
  
    // Printed after sleep is over
    fmt.Println("Sleep Over.....")
}


//https://stackoverflow.com/questions/31942163/how-to-write-my-own-sleep-function-using-just-time-after