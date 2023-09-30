package main

import (
	"log"
	//"uts/database"
	"uts/database"
	"uts/route"

	"github.com/gofiber/fiber/v2"
)

func main(){


	database.Connect()

	//Kumpulan Route Route 

	
	app.Post("/insert", route.InsertData )
	app.Get("/getData", route.GetAllData)
	app.Get("/getDataUser/:id_user", route.GetUserByid)


	app.Get("/delete/:id_user", route.Delete)
	app.Put("/update/:id_user", route.Update) 


	
	
}