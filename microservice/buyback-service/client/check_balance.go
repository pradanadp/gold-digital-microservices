package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pradanadp/gold-digital-microservices/buyback-service/helper"
	"github.com/pradanadp/gold-digital-microservices/buyback-service/model/web"
)

func CheckBalance(reqBody web.AccountRequest) web.WebResponse {
	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		log.Println("Error marshaling request body:", err)
		return web.WebResponse{}
	}

	req, err := http.NewRequest("GET", "http://localhost:3003/api/cek-saldo", bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		log.Println("Error creating request:", err)
		return web.WebResponse{}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return web.WebResponse{}
	}
	defer resp.Body.Close()

	var acc web.WebResponse
	helper.ReadFromResponseBody(resp, &acc)

	return acc
}
