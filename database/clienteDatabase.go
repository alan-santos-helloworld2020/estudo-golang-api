package database

import (
	"app/app/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



type Cliente struct{
		gorm.Model
		Nome string
		Email string
		Telefone string
		Cep string
}



func Conectar() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./banco/banco.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	db.AutoMigrate(&Cliente{})

	return db
}

func GetClientes() []Cliente{

	var clientes []Cliente
	Conectar().Find(&clientes);
    return clientes
}

func GetClienteById(id *int) Cliente{

	var cliente Cliente
	Conectar().First(&cliente, id);
    return cliente
}




func SalvarCliente(cl *models.ClienteDto)  {

	Conectar().Create(&Cliente{
		Nome: cl.Nome,
		Email: cl.Email,
		Telefone: cl.Telefone,
		Cep: cl.Cep,
	})
	
	
}