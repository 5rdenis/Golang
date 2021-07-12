// Реализовать композицию структуры Action
// от родительской структуры Human.
package main
  
import "fmt"
  

type Human struct {
    name string
    age  int
}
  

type Action struct {
    Human
    favSport  string
}
  

func (person Human) show() {
    fmt.Println("Name:",person.name,"Age:",person.age)
}
  

func (a Action) showAction() {
    fmt.Println(a.Human.name,"likes", a.favSport)
}
  
func main() {
  
    person := Human{name: "Denis", age: 22}
    actionPerson := Action{favSport: "volleyball", Human: person}
  
    actionPerson.show()
    actionPerson.showAction()
}