package client

import (
	"log"
	"net/http"

	"github.com/pradanadp/gold-digital-microservices/topup-service/helper"
	"github.com/pradanadp/gold-digital-microservices/topup-service/model/web"
)

func CheckPrice() web.WebResponse {
	req, err := http.NewRequest("GET", "http://localhost:3300/api/cek-harga", nil)
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

	var price web.WebResponse
	helper.ReadFromResponseBody(resp, &price)

	return price
}
