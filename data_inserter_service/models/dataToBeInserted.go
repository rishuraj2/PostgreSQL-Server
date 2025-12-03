package models

import "time"

type DataToBeInserted struct {
	StartTime           time.Time
	RawMaterialName     string
	RawMaterialQuantity int32
	Temperature         int32
	EndTime             time.Time
}
