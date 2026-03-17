package main

import (
	// "finance-help/config"
	// "fmt"
	"log"
	"net/http"
	// "strings"
)

func main() {
	// cfg := config.LoadConfig()
	// log.Println("Config Loaded")

	// client := &http.Client{}

	// TODO: get rid of tgbotapi dependency...
	// TODO: extract from ngrok the public endpoint and set as webhook
	// {{domain}}/bot{{token}}/setWebhook
	// url := fmt.Sprintf("%s/bot%s/setWebhook", cfg.TelegramApiEndpoint, cfg.TelegramToken)
	// body := fmt.Sprintf(`{"url": "%s"}`, "https://1d96-189-203-103-72.ngrok-free.app/test")
	// req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))

	// if err != nil {
	// 	log.Fatal("Err creating request: ", err)
	// }

	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal("Err sending request: ", err)
	// }

	// log.Println("Response status: ", resp.Status)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request")
		log.Println(r.Body) 
		w.WriteHeader(http.StatusOK)
 	    w.Write([]byte("OK"))
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Err starting server: ", err)
	}
}
