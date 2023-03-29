package main

import (
	  analyse "tweets/analyse"
	  db      "tweets/db"
)

func main() {
	
   dbConn, err := db.ConnectDB("localhost", "postgres", "postgres", "postgres", "5432")
   
   if err != nil {
   	  panic(err)
   }

   retVal, err := analyse.Sentiment(dbConn)

   if retVal != true || err != nil {
   	  panic(err)
   }
}
