package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "fmt"
  "strconv"
)


type Roll struct {
  ID            string `json:"id"`
  ImageNumber   string `json: "imageNumber"`
  Name          string `json: "name"`
  Ingredients   string `json: "ingredients"`
}

// init a slice of rolls
var rolls []Roll

func getRolls(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(rolls)

}
func getRoll(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  rollId := params["id"]

  for _, roll := range rolls {
    if roll.ID == rollId {
      json.NewEncoder(w).Encode(roll)
      return
    }
  }

  http.NotFound(w, r)
}

func createRoll(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  var newRoll Roll
  json.NewDecoder(r.Body).Decode(&newRoll)
  newRoll.ID = strconv.Itoa(len(rolls) + 1)

  rolls = append(rolls, newRoll)
  fmt.Println(rolls)

  json.NewEncoder(w).Encode(newRoll)
}

func updateRoll(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)

  for i, element := range(rolls) {
    if element.ID == params["id"] {
      // remove element from existing rolls slice by recreating it but omitting the found element's index
      rolls = append(rolls[:i], rolls[i+1:]...)
      var newRoll Roll
      json.NewDecoder(r.Body).Decode(&newRoll)
      newRoll.ID = params["id"]
      rolls = append(rolls, newRoll)
      json.NewEncoder(w).Encode(newRoll)
      return
    }
  }
}

func deleteRoll(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)

  for i, element := range(rolls) {
    if element.ID == params["id"] {
      // remove element from existing rolls slice by recreating it but omitting the found element's index
      rolls = append(rolls[:i], rolls[i+1:]...)
      json.NewEncoder(w).Encode(rolls)
      return
    }
  }
}

func main() {

  rolls = append(rolls, Roll{ID: "1", ImageNumber: "8", Name: "Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, Rice"})

  // init router
  router := mux.NewRouter()

  //endpoints
  router.HandleFunc("/sushi", getRolls).Methods("GET")
  router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
  router.HandleFunc("/sushi", createRoll).Methods("POST")
  router.HandleFunc("/sushi/{id}",updateRoll).Methods("PUT")
  router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":5000", router))
}
