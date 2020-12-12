package models

import (
	"time"
)

//Order - Modelo de persistência
type Order struct {
	Cpf                string
	Private            int
	Incompleto         int
	DataUltimaCompra   *time.Time
	TicketMedio        float64
	TicketUltimaCompra float64
	LojaMaisFrequente  string
	LojaUltimaCompra   string
}
