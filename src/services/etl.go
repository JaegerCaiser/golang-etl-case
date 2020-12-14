package services

import (
	"etl-neoway-challenge/src/database"
	"etl-neoway-challenge/src/models"
	"fmt"
	"time"
)

//Etl - Service de extração - tranformação e carregamento de dados.
func Etl() {
	start := time.Now()
	database.Migrate()

	eChan := make(chan []string)
	tChan := make(chan *models.Order)
	done := make(chan bool)

	go Extract(eChan)
	go Transform(eChan, tChan)
	go Load(tChan, done)

	<-done
	fmt.Println(time.Since(start))
}
