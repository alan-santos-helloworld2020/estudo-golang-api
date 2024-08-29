package models

type ClienteDto struct{	
	Nome string `json: "none"`
	Email string `json: "email"`
	Telefone string `json: "telefone"`
	Cep string `json: "cep"`
}