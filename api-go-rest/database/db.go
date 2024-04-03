package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func CriarConexao() {
	dsn := "root:teste@tcp(localhost:3306)/gorm?parseTime=true&tls=false"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

}
