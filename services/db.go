package services

import (
  // "fmt"
  "database/sql"
  "os"

  _ "github.com/skuttleman/gin-server/Godeps/_workspace/src/github.com/go-sql-driver/mysql"
  "github.com/markbates/goth"
)

type User struct {
  Id int
  Email string
  Name string
  Social_id string
  Image string
}

var db *sql.DB

func init() {
  db, _ = sql.Open("mysql", os.Getenv("DATABASE_URL"))
}

func UpdateOrCreate(u interface{}, err error) *User {
  user, err := getUser(u.(goth.User).Email)
  if err != nil {
    // fmt.Println(err)
    _, _ = insertUser(u.(goth.User))
    user, _ = getUser(u.(goth.User).Email)
  }
  // fmt.Println("db:" + user.email)
  return user
}

func getUser(email string) (*User , error) {
  query := "SELECT id, email, name, social_id, image FROM players WHERE email='" + email + "'"
  var user User
  err := db.QueryRow(query).Scan(&user.Id, &user.Email, &user.Name, &user.Social_id, &user.Image)
  if err != nil {
    return nil, err
  } else {
    return &user, nil
  }
}

func insertUser(user goth.User) (*sql.Rows, error) {
  query := "INSERT INTO players (email, name, social_id, image) VALUES ('" +
    user.Email + "', '" + user.Name + "', '" + user.UserID + "', '" + user.AvatarURL + "');"
  return db.Query(query)
}
