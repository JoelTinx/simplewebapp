package main

import (
	"log"
	"net/http"
	"handlers"
)

func main()  {
	fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/",          handlers.IndexController)
  http.HandleFunc("/login",     handlers.LoginController)
  http.HandleFunc("/admin", 	handlers.ManagerController)

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  log.Printf("Listening in port: %s", port)
  http.ListenAndServe(":" + port, nil)
}
