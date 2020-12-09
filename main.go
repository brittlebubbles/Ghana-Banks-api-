package main

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "io/ioutil"
  "log"
  "net/http"
)


// Structs
type Bank struct {
  Name string `json:"name"`
  Address string `json:"address"`
  PhoneNumber string `json:"phoneNumber"`
  Website string `json:"website"`
  Email string `json:"email"`
}

type BankJSON struct {
  Name string `json:"name"`
  Address string `json:"address"`
  PhoneNumber string `json:"phoneNumber"`
  Website string `json:"website"`
  Email string `json:"email"`
}

var banks []BankJSON

func getAllBanks(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")

  var newBanks []Bank

  for _, bank := range banks {
    newBanks = append(newBanks, Bank{
      Name: bank.Name,
      Address: bank.Address,
      PhoneNumber: bank.PhoneNumber,
      Website: bank.Website,
      Email: bank.Email,
    })
  }

  _ = json.NewEncoder(w).Encode(newBanks)
}

func main() {

  bankData, err := ioutil.ReadFile("./bankdata.json")

  if err != nil {
    log.Fatal(err)
  }

  if err := json.Unmarshal(bankData, &banks); err != nil {
    log.Fatal(err)
  }

  r := mux.NewRouter()

  //Initialize Routes
  r.HandleFunc("/banks", getAllBanks).Methods("GET")

  // Start server
  log.Fatal(http.ListenAndServe(":8000", r))
}
