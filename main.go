package main

import (
	"CurrencyNotifications/Tokens"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", Tokens.TgToken))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
