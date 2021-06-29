package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "strconv"

    "github.com/gorilla/mux"
)



func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome!\n")
}

func AccountIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(Accounts); err != nil {
        panic(err)
    }
}

func AccountShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    accountId, err := strconv.Atoi(vars["accountId"])
    if err != nil {
        panic(err)
    }
    //t := RepoFindAccount(accountId)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(Accounts[accountId]); err != nil {
        panic(err)
    }
}

func AccountCreate(w http.ResponseWriter, r *http.Request) {
    var account Account
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &account); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
 
    t := RepoCreateAccount(account)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}

func AccountUpdate(w http.ResponseWriter, r *http.Request) {
    var account Account
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &account); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
 
    t := RepoUpdateAccount(account.Id,account)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}


func TransferMoney (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    var transfer Transfer
    _ = json.NewDecoder(r.Body).Decode(&transfer)
    idFrom := transfer.IdFrom
    idTo := transfer.IdTo
    if *transfer.Value <= *Accounts[idFrom].Balance {
        *Accounts[idFrom].Balance -= *transfer.Value
        *Accounts[idTo].Balance += *transfer.Value
        w.WriteHeader(http.StatusOK)
    }
}


/*
func updateUser(c echo.Context) error {
    u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    id, _ := strconv.Atoi(c.Param("id"))
    users[id].Name = u.Name
    return c.JSON(http.StatusOK, users[id])
}
*/
