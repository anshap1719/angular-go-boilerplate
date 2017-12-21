package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "os"
  "log"
  "angular-go-boilerplate/src/server/utils"
  "fmt"
  "github.com/rs/cors"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/hello-world", helloWorld)

  // Solves Cross Origin Access Issue
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:4200"},
  })
  handler := c.Handler(r)

  srv := &http.Server{
    Handler: handler,
    Addr:    ":" + os.Getenv("PORT"),
  }

  log.Fatal(srv.ListenAndServe())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
  var data = struct {
    Title string `json:"title"`
  }{
    Title: "Golang + Angular Starter Kit",
  }

  jsonBytes, err := utils.StructToJson(data); if err != nil {
    fmt.Print(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonBytes)
  return
}
