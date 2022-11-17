package main

import (
	"database/sql"
	"log"
	//"bufio"
	"fmt"

	"github.com/joho/godotenv"

	//"encoding/json"
	db "github.com/avito-tech/database"
	"io"
	"os"
)

type client struct{
	ID string `json: "id"`
	AccBalance int `json: "accbalance"`
}

// func refill(id string, sum int)  MarshalJSON() ([]byte, error) {

// 	//we have already checked the database and if this client is new we've add them
// 	//we need to sum this income to theirs current balance in our database sheet
// 	return json.Marshal()
// }

// func check_if_new(id string){
// 	//checking the database on having this id
// 	//if not adding a new client with 0 on theirs balance
// 	NewClient:=new(client)
// }

// func check(id string) MarshalJSON() ([]byte, error) {
// 	//showing acc balance of this client
// 	return json.Marshal()
// }

func main(){

	envErr := godotenv.Load(); if envErr != nil {
		fmt.Printf("Error loading credentials: %v", envErr)
	 }

	 
	 var (
		password = os.Getenv("MSSQL_DB_PASSWORD")
		user = os.Getenv("MSSQL_DB_USER")
		port = os.Getenv("MSSQL_DB_PORT")
		database = os.Getenv("MSSQL_DB_DATABASE")
	 )

	 ConnectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", user, password, port, database)

	 sqlObj, connectionError := sql.Open("mssql", database.ConnectionString); if connectionError != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
	 }
  
	 data := database.Database{
		SqlDb: sqlObj,
	 }
   
	const jsonStream = `
	{"header": "refill", "id": "908887789988", "income": 10000}
	`


	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		h:= m.header
	
		switch h {
		case "refill":
			db.Check_if_new(m.id)
			db.Refill(m.id, m.income)
		case "balanse check":
			db.Check_if_new(m.id)
			db.Balance_check(m.id)

		}
	}
}