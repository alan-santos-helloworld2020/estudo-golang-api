package models


type UserDto struct{
	Username string `uri:"username"`
	Password string `uri:"password"`
}