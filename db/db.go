package db

import (
  "database/sql"
  "fmt"
_ "github.com/lib/pq"
  tinfo "tweets/json"
)

var err error

func ConnectDB(host string, dbname string, username string, password string,  port string) (*sql.DB, error) {

   psqlInfo := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s sslmode=disable", 
                           host, dbname, username, password, port)

   
   db, _ := sql.Open("postgres", psqlInfo)

   
   //defer db.Close()

   err := db.Ping()

   if err != nil {
      return nil, err
   }

   return db, nil

}

func InsertDB(db *sql.DB, file tinfo.TweetInfo) (bool, error) {
     

   for _, tweet:= range file.Tweet {

      sqlStatement := `INSERT INTO my_schema.analyser (id, text)
                       VALUES ($1, $2)`

      _, err = db.Exec(sqlStatement, tweet.Id, tweet.Text)
   
      if err != nil {
        return false, err
      }

   }

   return true, nil

}
