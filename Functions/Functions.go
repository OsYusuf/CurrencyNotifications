package Functions

import (
	"CurrencyNotifications/Structs"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Polling(TgToken string) (string, int) {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=-1", TgToken))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	var req Structs.Req
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Fatal(err)
	}
	return req.Result[0].Message.Text, req.Result[0].Message.From.ID
}

func SendMessage(TgToken string, message string, chatId int) {
	_, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", TgToken, chatId, message))
	if err != nil {
		log.Fatal(err)
	}
}
