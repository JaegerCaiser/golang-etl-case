package src

import (
	"etl-neoway-challenge/src/models"
	"etl-neoway-challenge/src/utils"
	"strconv"
)

//Transform - Transforma valores do array de string para struct
func Transform(eChan chan []string, tChan chan *models.Order) {
	for record := range eChan {
		o := new(models.Order)
		o.Cpf = utils.ConvertOnlyNumber(record[0])
		o.Private, _ = strconv.Atoi(record[1])
		o.Incompleto, _ = strconv.Atoi(record[2])
		o.DataUltimaCompra = utils.ConvertStringToDate(record[3])
		o.TicketMedio, _ = strconv.ParseFloat(record[4], 64)
		o.TicketUltimaCompra, _ = strconv.ParseFloat(record[5], 64)
		o.LojaMaisFrequente = utils.ConvertOnlyNumber(record[6])
		o.LojaUltimaCompra = utils.ConvertOnlyNumber(record[7])
		tChan <- o
	}
	close(tChan)
}
