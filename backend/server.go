package main

import (
  "log"
  "net/http"
  "simplewebapp/handlers"
  "simplewebapp/routes"
)

func main() {
  fs := http.FileServer(http.Dir("public"))
  // http.Handle("/public/", http.StripPrefix("/public/", fs))
  http.Handle(routes.IndexRoute())

  // Get port of enviroment variables
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  http.ListenAndServe(":" + port)
}
