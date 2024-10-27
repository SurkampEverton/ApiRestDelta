package database

import (
	"log"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=177.72.162.73 user=cliente password=Hs6KnuJU dbname=Delta port=7000 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NoLowerCase:   false,
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		}})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
