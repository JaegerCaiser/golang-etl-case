package src

import (
	"etl-neoway-challenge/src/models"
	"fmt"
	"os"
	"sync"
)

//Load - Carrega os dados no arquivo
func Load(tChan chan *models.Order, done chan bool) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()

	fmt.Fprintf(f, "%20s%16s%13s%13s%16s%16s%20s%20s", "Cpf", "Private",
		"Incompleto", "Data da Ultima Compra", "Ticket Medio", "Ticket da Ultima Compra", "Loja Mais Frequente", "Loja da Ultima Compra\n")

	var wg sync.WaitGroup
	for o := range tChan {
		wg.Add(1)
		go func(o *models.Order) {
			// time.Sleep(1 * time.Millisecond)
			fmt.Fprintf(f, "%20s %15d %d %s %15.2f%15.2f %20s %20s \n",
				o.Cpf, o.Private, o.Incompleto, o.DataUltimaCompra,
				o.TicketMedio, o.TicketUltimaCompra, o.LojaMaisFrequente, o.LojaUltimaCompra)
			defer wg.Done()
		}(o)
	}

	wg.Wait()
	done <- true
}
