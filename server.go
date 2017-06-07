package main

import (
	"log"
	"net/http"
	"controllers"
)

func main()  {
	fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/",          controllers.IndexController)

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  log.Printf("Listening in port: %s", port)
  http.ListenAndServe(":" + port, nil)
}