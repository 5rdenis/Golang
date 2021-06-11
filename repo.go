package main
 
import "fmt"
 
var currentId int
 
var accounts Accounts
 
// Give us some seed data
func init() {
    RepoCreateAccount(Account{Name: "Denis",Balance: 700})
    RepoCreateAccount(Account{Name: "Max"})
}
 
func RepoFindAccount(id int) Account {
    for _, a := range accounts {
        if a.Id == id {
            return a
        }
    }
    // return empty Account if not found
    return Account{}
}
 
func RepoCreateAccount(a Account) Account {
    currentId += 1
    a.Id = currentId
    accounts = append(accounts, a)
    return a
}
 
func RepoDestroyAccount(id int) error {
    for i, a := range accounts {
        if a.Id == id {
            accounts = append(accounts[:i], accounts[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Account with id of %d to delete", id)
}