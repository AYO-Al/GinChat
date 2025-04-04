package main

import (
	"GinChat/router"
)

func main() {
	//err := app.Db.AutoMigrate(&models.UserBasic{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	r := router.Router()
	r.Run()

}
