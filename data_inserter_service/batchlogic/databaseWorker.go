package batchlogic

import (
	"data_inserter_service/models"
	"database/sql"
	"log"
	//"time" //add this when slow consumer is to be used

	_ "github.com/lib/pq"
)

func DatabaseWorker(db *sql.DB, insertChan <-chan models.DataToBeInserted) {
	for data := range insertChan {
		insertBatchData(db, data)
	}
	//time.Sleep(1 * time.Second); //use it to simulate slow consumer
}

func insertBatchData(db *sql.DB, data models.DataToBeInserted) {
	_, err := db.Exec(`
	INSERT INTO batch_inputs(start_time, raw_material_name, raw_material_quantity, temperature, end_time)
	VALUES ($1, $2, $3, $4, $5)
	`, data.StartTime, data.RawMaterialName, data.RawMaterialQuantity, data.Temperature, data.EndTime)

	if err != nil {
		log.Fatal("Error inserting data: ", err)
	} else {
		log.Printf("Inserted Batch Data: raw_material_name: %s, raw_material_quantity: %d | %s -> %s", data.RawMaterialName, data.RawMaterialQuantity, data.StartTime.Format("15:04:05"), data.EndTime.Format("15:04:05"))
	}
}