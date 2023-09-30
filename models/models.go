package models


type User struct {
	Id_user uint `json:"id_user" gorm:"primaryKey;autoIncrement:true"`
	Nama, Email , Username string
	Password string
} 

//Tulis jawab code nya di sini 