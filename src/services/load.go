package services

import (
	"etl-neoway-challenge/src/daos/daos"
	"etl-neoway-challenge/src/models"
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
			go daos.SaveList(arrayOrders, &wg)
			arrayOrders = make([]*models.Order, 0, bufferSize)
		}
	}

	//Salvando os remanescentes
	if len(arrayOrders) > 0 {
		wg.Add(1)
		go daos.SaveList(arrayOrders, &wg)
	}

	wg.Wait()
	done <- true
}
