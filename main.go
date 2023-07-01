package main

import (
	"CurrencyNotifications/Structs"
	"CurrencyNotifications/Tokens"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Polling(TgToken string) string {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", TgToken))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	var req Structs.Req
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}
	return req.Result[0].Message.Text
}

func main() {
	fmt.Println(Polling(Tokens.TgToken))
}
