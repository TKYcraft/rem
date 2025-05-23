package main

import (
	"rem/bot"
	"rem/db"
)

func main() {
	db.InitDB()
	bot.Start()
}