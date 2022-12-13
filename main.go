package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type UserAdmin struct {
    name    string `bson:"first_name,omitempty"`
    phone   string `bson:"phone,omitempty"`
    address string `bson:"address,omitempty"`
    cvv     string `bson:"cvv,omitempty"`
    id      string `bson:"id,omitempty"`
}

func getUser() UserAdmin {
    u := UserAdmin{
        name:    "jo",
        phone:   "6343",
        address: "aaaa",
        cvv:     "666",
    }

    return u
}

//func homeLink(w http.ResponseWriter, r *http.Request) {
//    getUser()
//
//    fmt.Fprintf(w, "Welcome home!")
//}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    //router.HandleFunc("/api/", homeLink).Methods("GET")
    router.HandleFunc("/api/post", func(w http.ResponseWriter, r *http.Request) {
        getUser()
    }).Methods("GET")

    log.Fatal(http.ListenAndServe(":4356", router))
}
