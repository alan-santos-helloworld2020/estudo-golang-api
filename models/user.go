package models


type User struct{
	Id int `uri:"id"`
	Username string `uri:"username"`
	Password string `uri:"password"`
}
