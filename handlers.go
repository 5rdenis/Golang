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
    if err := json.NewEncoder(w).Encode(accounts); err != nil {
        panic(err)
    }
}

func AccountShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    accountId, err := strconv.Atoi(vars["accountId"])
    if err != nil {
        panic(err)
    }
    t := RepoFindAccount(accountId)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t.Balance); err != nil {
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
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    vars := mux.Vars(r)
    accountId, err := strconv.Atoi(vars["accountId"])
    if err != nil {
        panic(err)
    }
    for index, item := range accounts {
        if item.Id == accountId {
            accounts = append(accounts[:index], accounts[index+1:]...)
            var account Account
            _ = json.NewDecoder(r.Body).Decode(&account)
            return
        }
    }
    json.NewEncoder(w).Encode(accounts)
}

