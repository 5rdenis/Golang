//Написать программу, которая переворачивает слова в строке (snow dog sun - sun dog snow).
package main
 
import (
   "fmt"
   "strings"
)
 
func main() {
   str := "snow dog sun yikes dog snow snow sun sun lol lmao fgk"
   words := strings.Split(str, " ")
   //Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
   len := len(words)
   for i := 0; i<len/2 ; i++ {
      //tmp := words[len-i-1]
      //words[len-i-1] = words[i]
      words[len-i-1], words[i] = words[i],words[len-i-1]
      //words[i] = tmp
   }
   str = strings.Join(words," ")
   fmt.Println(str)
}