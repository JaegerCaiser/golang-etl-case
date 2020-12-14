package database

import "etl-neoway-challenge/src/models"

//Migrate - Cria estrutura na base de dados
func Migrate() {
	Connect().AutoMigrate(&models.Order{})
}
