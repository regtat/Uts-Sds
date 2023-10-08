package models

type User struct {
	Id_user               uint `json:"id_user" gorm:"primaryKey;autoIncrement:true"`
	Nama, Email, Username string
	Password              string
}

// Tulis jawab code nya di sini
type Admin struct {
	Id_admin   uint `json:"id_admin" gorm:"primaryKey;autoIncrement:true"`
	Nama_admin string
	Email      string
	Password   string
}
