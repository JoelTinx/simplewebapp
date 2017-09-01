package main

import (
	"log"
	"net/http"
	"handlers"
)

func init()  {
	log.Println("Inicializando servicios base")
}

const (

)

func main()  {
	fs := http.FileServer(http.Dir("public"))
  http.Handle("/public/", http.StripPrefix("/public/", fs))

  http.HandleFunc("/",          handlers.IndexController)
  http.HandleFunc("/login",     handlers.LoginController)
  http.HandleFunc("/admin", 	handlers.ManagerController)


	// Session required path's
	http.HandleFunc("/api/v1/path01", 	handlers.ManagerController)
	http.HandleFunc("/api/v1/path02", 	handlers.ManagerController)
	http.HandleFunc("/api/v1/path03", 	handlers.ManagerController)


  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  log.Printf("Listening in port: %s", port)
  http.ListenAndServe(":" + port, nil)
}

/*
Configurar variables de entorno desde la terminal
*/
