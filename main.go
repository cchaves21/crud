package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "mysecretpassword"
  dbname   = "test"
)

func main() {
 
  //insert(31, "kchaves@gmail.com", "Carlos", "chaves")
  //update(1, "James", "Oliver")
  //delete(1)
    
}

func dbConn() (db *sql.DB) {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  return db
}

func insert(age int, email string, first_name string, last_name string) {
  db := dbConn()

  sqlStatement := `
  INSERT INTO users (age, email, first_name, last_name)
  VALUES ($1, $2, $3, $4)
  RETURNING id`

  var err error
  id := 0

  err = db.QueryRow(sqlStatement, age, email, first_name, last_name).Scan(&id)

  if err != nil {
    panic(err)
  }

  defer db.Close()
  fmt.Println("New record ID is:", id)
}

func update(id int, NewFirst string, NewLast string ) {
  db := dbConn()

  sqlStatement := `
  UPDATE users
  SET first_name = $2, last_name = $3
  WHERE id = $1;`

  _, err := db.Exec(sqlStatement, id, NewFirst, NewLast)

  if err != nil {
    panic(err)
  }
  defer db.Close()
}

func delete(id int) {
  db := dbConn()

  sqlStatement := `
  DELETE FROM users
  WHERE id = $1;`

  _, err := db.Exec(sqlStatement, id)

  if err != nil {
    panic(err)
  }
  defer db.Close()
}