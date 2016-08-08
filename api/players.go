package api

import (
  "net/http"
  "fmt"
  "os"

  "github.com/gorilla/sessions"
  // "github.com/skuttleman/capstone/services"
  "github.com/gorilla/pat"
)

func Players(prefix string, app *pat.Router) {
  store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

  // All
  app.Get(prefix + "/current-user", func(res http.ResponseWriter, req *http.Request) {
    session, _ := store.Get(req, "session")
    user := session.Values["user"]
    fmt.Println(user)
    fmt.Fprint(res, user)
  })

  app.Get(prefix, func(res http.ResponseWriter, req *http.Request) {
    // List players, group by ever played with, order by recently played with
    fmt.Fprint(res, map[string]interface{}{
      "test": true,
    })
  })



}
