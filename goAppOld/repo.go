package main
 
var currentId int
var Accounts = map[int]*Account{}



func RepoCreateAccount(a Account) Account {
    currentId += 1
    Accounts[currentId] = &a
    Accounts[currentId].Id = currentId
    return a
}

func RepoUpdateAccount(id int, a Account) Account {
	if (Accounts[id].Name != a.Name && a.Name != "") {
		Accounts[id].Name = a.Name
	} 
	if (Accounts[id].Balance != a.Balance && a.Balance != nil) {
		Accounts[id].Balance = a.Balance
    }
    return a

}