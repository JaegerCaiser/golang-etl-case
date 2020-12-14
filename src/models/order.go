package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Order - Modelo de persistência
type Order struct {
	gorm.Model
	Cpf                string
	DocucmentoValido   bool
	Private            int
	Incompleto         int
	DataUltimaCompra   *time.Time
	TicketMedio        float64
	TicketUltimaCompra float64
	LojaMaisFrequente  string
	LojaUltimaCompra   string
}
