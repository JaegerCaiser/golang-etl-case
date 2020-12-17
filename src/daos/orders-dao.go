package daos

import (
	"etl-neoway-challenge/src/database"
	"etl-neoway-challenge/src/models"
	"log"
	"sync"
)

const bufferSize = 1000

//SaveList - Salva lista na base de dados - GOROUTINE
func SaveList(orders []*models.Order, wg *sync.WaitGroup) {
	tx := database.Connect().CreateInBatches(orders, bufferSize)
	log.Println(tx.RowsAffected, "Registros salvos com sucesso")
	defer wg.Done()
}
