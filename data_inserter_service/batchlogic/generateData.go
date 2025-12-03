package batchlogic

import (
	"data_inserter_service/models"
	"math/rand"
	"time"
)

var (
	rawMaterialNames = []string{"water", "caustic_soda", "perfume", "glycerin"}
)

func GenerateData(sequence *int, timeKeeper *time.Time) models.DataToBeInserted {
	index := *sequence % len(rawMaterialNames)
	data := models.DataToBeInserted{
		StartTime:           *timeKeeper,
		RawMaterialName:     rawMaterialNames[index],
		RawMaterialQuantity: getRawMaterialQuantity(rawMaterialNames[index]),
		Temperature:         20 + rand.Int31n(20),
		EndTime:             (*timeKeeper).Add(2 * time.Minute),
	}
	(*sequence)++
	*timeKeeper = data.EndTime
	return data
}

func getRawMaterialQuantity(rawMaterialName string) int32 {
	switch rawMaterialName {
	case "water":
		return 40 + rand.Int31n(20)

	case "caustic_soda":
		return 8 + rand.Int31n(5)

	case "perfume":
		return 2 + rand.Int31n(2)

	case "glycerin":
		return 50 + rand.Int31n(20)

	default:
		return 0
	}
}
