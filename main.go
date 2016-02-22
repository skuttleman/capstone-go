package main

import (
  "os"
  "net/http"
  // "fmt"

  "github.com/gorilla/pat"
  // "github.com/gorilla/sessions"
  "github.com/gorilla/context"
  _ "github.com/joho/godotenv/autoload"
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
  // store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

  app.Get("/", func(res http.ResponseWriter, req *http.Request) {
    file := req.URL.Path[1:]
    if file == "" {
      file = "index.html"
    }
    http.ServeFile(res, req, "./front-end/public/" + file)
    // fmt.Fprintf(res, "Hi there, I love %s!", req.URL.Path[1:])
  })
  http.ListenAndServe(":" + port, context.ClearHandler(app))
}
