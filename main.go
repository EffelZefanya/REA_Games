package main

import (
	"fmt"
	"rea_games/config"
)

func main(){
	db := config.ConnectDB()
	fmt.Println(db)
}