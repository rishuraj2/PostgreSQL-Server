package main

import (
	"data_inserter_service/batchlogic"
	"data_inserter_service/models"
	"data_inserter_service/utility"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	schemaType = "public"
	tableName  = "batch_inputs"
	timeKeeper = time.Now()
	sequence   = 0
)

func main() {
	db, err := utility.ConnectDB()

	if err != nil {
		log.Fatal("Error Connection to database: ", err)
	}
	defer db.Close()

	if err := utility.CheckDbConnection(db); err != nil {
		log.Fatal("Failed to connect to database: ", err)
		return
	}

	if tableExists := utility.CheckTableExists(db, schemaType, tableName); tableExists {
		log.Println("Table found and ready!")
	} else {
		log.Fatalf("Failed to find table %s! Check 'Scheme Type' and Table Name'", tableName)
	}

	insertChan := make(chan models.DataToBeInserted, 100)
	go batchlogic.DatabaseWorker(db, insertChan)

	for {
		insertChan <- batchlogic.GenerateData(&sequence, &timeKeeper)
		time.Sleep(1 * time.Second)
	}
}
