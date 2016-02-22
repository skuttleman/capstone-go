package services

import (
  "fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/pat"
  "github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/gplus"
)

func Auth() {
  goth.UseProviders(
		gplus.New(os.Getenv("GOOGLE_CLIENT_ID"), os.Getenv("GOOGLE_CLIENT_SECRET"), "http://localhost:8000/auth/gplus/callback"),
	)

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {
		// print our state string to the console. Ideally, you should verify
		// that it's the same string as the one you set in `setState`
		// fmt.Println("State: ", gothic.GetState(req))

		user, err := gothic.CompleteUserAuth(res, req)

		if err != nil {
			// fmt.Fprintln(res, err)
			return
		}

		session, err := store.Get(req, "session")
		session.Values["user"] = user
		session.Save(req, res)
		fmt.Println(session.Values["user"])

		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, user)
	})

	p.Get("/auth/logout", func(res http.ResponseWriter, req *http.Request) {
		session, _ := store.Get(req, "session")
		session.Values["user"] = nil
		session.Save(req, res)
	})
	p.Get("/auth/{provider}", gothic.BeginAuthHandler)
	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		session, _ := store.Get(req, "session")
		fmt.Println(session.Values["user"])
		t, _ := template.New("foo").Parse(indexTemplate)
		t.Execute(res, session.Values["user"])
	})
}
