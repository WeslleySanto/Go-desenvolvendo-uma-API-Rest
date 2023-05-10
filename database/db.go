package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Conecta() {
	dns := "./personalidades.db"
	DB, err = gorm.Open(sqlite.Open(dns))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
