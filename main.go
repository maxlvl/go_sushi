package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)


type Roll struct {
  ID            string `json:"id"`
  ImageNumber   string `json: "imageNumber"`
  Name          string `json: "name"`
  Ingredients   string `json: "ingredients"`
}

// init a slice of rolls
var rolls []Roll

func main() {

  // init router
  router := mux.NewRouter()

  //endpoints
  router.Handlefunc("/sushi", getRolls).Methods("GET")
  router.Handlefunc("/sushi/{id}", getRoll).Methods("GET")
  router.Handlefunc("/sushi", createRoll).Methods("POST")
  router.Handlefunc("/sushi/{id}",updateRoll).Methods("PUT")
  router.Handlefunc("/sushi/{id}", deleteRoll).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":5000", router))
}
