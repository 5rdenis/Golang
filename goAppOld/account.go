package main
 
type Account struct {
    Id        int       `json:"id"`
    Name      string    `json:"name"`
    Balance   *int      `json:"balance"`
}

type Transfer struct {
    IdFrom        int       `json:"idfrom"`
    IdTo          int       `json:"idto"`
    Value         *int       `json:"value"`
}