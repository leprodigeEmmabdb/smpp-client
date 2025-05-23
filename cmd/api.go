package main

import (
	"encoding/json"
	"net/http"
	"github.com/tonpseudo/smpp-client/internal/smppclient"
)

type SendRequest struct {
	Message string   `json:"message"`
	Numbers []string `json:"numbers"`
}

func main() {
	client := smppclient.NewClient()
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		var req SendRequest
		json.NewDecoder(r.Body).Decode(&req)

		go smppclient.SendMultiple(client.Transceiver, req.Numbers, req.Message)
		w.WriteHeader(http.StatusAccepted)
	})

	http.ListenAndServe(":8080", nil)
}
