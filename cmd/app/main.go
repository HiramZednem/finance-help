package main

import (
	"bufio"
	"finance-help/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	cfg := config.LoadConfig()
	log.Println("Config Loaded")

	client := &http.Client{}

	// TODO: get rid of tgbotapi dependency...
	// TODO: extract from ngrok the public endpoint and set as webhook
	// {{domain}}/bot{{token}}/setWebhook
	url := fmt.Sprintf("%s/bot%s/setWebhook", cfg.TelegramApiEndpoint, cfg.TelegramToken)

	if cfg.ENV == "dev" {
		urlChan := make(chan string)

		cmd := exec.Command("ngrok", "http", "8080", "--log=stdout")

		reader, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal("Error creando el pipe:", err)
		}

		if err := cmd.Start(); err != nil {
			log.Fatal("Error iniciando ngrok:", err)
		}

		go func() {
			sc := bufio.NewScanner(reader)
			for sc.Scan() {
				if strings.Contains(sc.Text(), "url=") {
					_, after, _:= strings.Cut(sc.Text(), "url=");
					urlChan <- after
				}
			}
		}()

		publicURL := <-urlChan
		log.Println(fmt.Sprintf("URL NGROK: %s", publicURL))

		body := fmt.Sprintf(`{"url": "%s"}`, )
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))

		if err != nil {
			log.Fatal("Err creating request: ", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Err sending request: ", err)
		}

		log.Println("Response status: ", resp.Status)
	}

	

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request")

		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Err reading body: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		log.Println("Request body: ", string(bodyBytes))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Err starting server: ", err)
	}
}
