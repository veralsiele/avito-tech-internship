package database

import (
	"fmt"
	"database/sql"
 )
 
 func (db Database) refill(id string, sum int)  MarshalJSON() ([]byte, error) {

	//we have already checked the database and if this client is new we've add them
	//we need to sum this income to theirs current balance in our database sheet
	
	err = db.SqlDb.PingContext(dbContext); if err != nil {
		return -1, err
	  }

	
	return json.Marshal()
}

func check_if_new(id string){
	//checking the database on having this id
	//if not adding a new client with 0 on theirs balance
	
   var err error

   err = db.SqlDb.PingContext(dbContext); if err != nil {
      return -1, err
    }
	
   query, err := db.SqlDb.Prepare(queryStatement); if err != nil {
	return -1, err
    }

    defer query.Close()

	newRecord := query.QueryRowContext(dbContext,
		sql.Named("ID", id),
		sql.Named("Balance", 0),
	 )
  

}

func check(id string) MarshalJSON() ([]byte, error) {
	//showing acc balance of this client
	return json.Marshal()
}