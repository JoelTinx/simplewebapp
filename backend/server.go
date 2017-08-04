package main

import (
  "log"
  "net/http"
  "simplewebapp/handlers"
)

func main() {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  http.ListenAndServe(":" + port)
}
