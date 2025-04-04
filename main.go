package main

import "GinChat/router"

func main() {
	//db.Db.Debug().AutoMigrate(&models.UserBasic{})
	r := router.Router()
	r.Run()

}
