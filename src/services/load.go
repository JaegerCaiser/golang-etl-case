package services

import (
	"etl-neoway-challenge/src/database"
	"etl-neoway-challenge/src/models"
	"log"
	"sync"
)

const bufferSize = 1000

//Load - Carrega os dados no arquivo
func Load(tChan chan *models.Order, done chan bool) {
	arrayOrders := make([]*models.Order, 0, bufferSize)
	var wg sync.WaitGroup
	for o := range tChan {
		// time.Sleep(10 * time.Millisecond)
		if len(arrayOrders) < bufferSize {
			arrayOrders = append(arrayOrders, o)
		} else {
			wg.Add(1)
			go func(o []*models.Order) {
				tx := database.Connect().CreateInBatches(o, bufferSize)
				log.Println(tx.RowsAffected, "Registros salvos com sucesso")
				defer wg.Done()
			}(arrayOrders)
			arrayOrders = make([]*models.Order, 0, bufferSize)
		}
	}

	wg.Wait()
	done <- true
}
