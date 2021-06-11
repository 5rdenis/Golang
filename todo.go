package main
 
type Account struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Balance   int      `json:"balance"`
}
 
type Accounts []Account