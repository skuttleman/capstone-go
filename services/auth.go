package services

import (
  "fmt"
	// "html/template"
	"net/http"
	"os"

	"github.com/gorilla/pat"
  "github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/gplus"
)

var host string

func init() {
  host = os.Getenv("HOST")
  if host == "" {
    host = "http://localhost:8000"
  }
}

func Auth(prefix string, app *pat.Router) {
  goth.UseProviders(
		gplus.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), host + "/auth/gplus/callback"),
	)

	app.Get(prefix + "/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		user := UpdateOrCreate(gothic.CompleteUserAuth(res, req))
		cookie := userCookie(user)
		http.SetCookie(res, cookie)
		fmt.Fprint(res, "user: " + cookie.Value)
	})

	app.Get(prefix + "/logout", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("TODO: Clear cookie")
	})
	app.Get(prefix + "/{provider}", gothic.BeginAuthHandler)
}

func userCookie(user *User) *http.Cookie {
	var cookie http.Cookie
	cookie.Name = "user"
	cookie.Value = user.Email
	cookie.MaxAge = 60 * 60 * 24 * 30
	cookie.Secure = true
	return &cookie
}
