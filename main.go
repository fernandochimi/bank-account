package main

import (
    "github.com/gorilla/mux"
    "github.com/fernandochimi/bank-account/src/app"
    "github.com/fernandochimi/bank-account/src/controllers"
    "os"
    "fmt"
    "net/http"
)

func main() {

    router := mux.NewRouter()
    router.Use(app.JwtAuth)

    router.HandleFunc("/user/new", controllers.CreateAccount).Methods("POST")
    router.HandleFunc("/user/login", controllers.Authenticate).Methods("POST")
    router.HandleFunc("/operations", controllers.GetOperations).Methods("GET")
    router.HandleFunc("/operations/{id}", controllers.GetOperation).Methods("GET")
    router.HandleFunc("/account/{id}", controllers.GetAccount).Methods("GET")
    router.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    fmt.Println(port)

    err := http.ListenAndServe(":" + port, router)
    if err != nil {
        fmt.Print(err)
    }
}