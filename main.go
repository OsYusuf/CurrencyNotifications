package main

import (
	"CurrencyNotifications/Functions"
	"CurrencyNotifications/Tokens"
	"fmt"
	"time"
)

func main() {
	for {
		msg, id := Functions.Polling(Tokens.TgToken)
		fmt.Println(msg)
		fmt.Println(id)
		if msg == "/set" {
			Functions.SendMessage(Tokens.TgToken, "Put+text+of+message+here", id)
		}
		time.Sleep(time.Second)
	}
}
