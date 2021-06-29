package main
 
import "sync"

type Account struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Balance   *int      `json:"balance"`
}

type Transfer struct {
	Rw sync.RWMutex
    IdFrom        int       `json:"idfrom"`
    IdTo          int       `json:"idto"`
    Value         *int       `json:"value"`
}