package src

import (
	"etl-neoway-challenge/src/database"
	"etl-neoway-challenge/src/models"
	"sync"
	"time"
)

//Load - Carrega os dados no arquivo
func Load(tChan chan *models.Order, done chan bool) {
	var wg sync.WaitGroup
	for o := range tChan {
		time.Sleep(10 * time.Millisecond)
		wg.Add(1)
		go func(o *models.Order) {
			database.Connect().Create(o)
			defer wg.Done()
		}(o)
	}

	wg.Wait()
	done <- true
}
