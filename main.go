package main

import (
  _ "github.com/joho/godotenv/autoload"

  "os"
  "net/http"
  "fmt"
  // "reflect"

  "github.com/gorilla/pat"
  // "github.com/gorilla/context"
  // "github.com/markbates/goth"
	// "github.com/markbates/goth/gothic"
	// "github.com/markbates/goth/providers/gplus"

  "github.com/skuttleman/capstone/api"
  "github.com/skuttleman/capstone/services"
)

var port string

func init() {
  port = os.Getenv("PORT")
  if port == "" {
    port = "8000"
  }
}

func main() {
  app := pat.New()


  services.Auth("/auth", app)

  api.Players("/api/players", app)

  app.Get("/", func(res http.ResponseWriter, req *http.Request) {
    file := req.URL.Path[1:]
    if file == "" {
      file = "index.html"
    }
    http.ServeFile(res, req, "./front-end/public/" + file)
  })

  fmt.Println("Server is listening on port: " + port)
  http.ListenAndServe(":" + port, context.ClearHandler(app))
}
