package main

import (
	"etl-neoway-challenge/src"
	"etl-neoway-challenge/src/models"
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	eChan := make(chan []string)
	tChan := make(chan *models.Order)
	done := make(chan bool)

	go src.Extract(eChan)
	go src.Transform(eChan, tChan)
	go src.Load(tChan, done)

	<-done
	fmt.Println(time.Since(start))

}
