package main

import (
	"bufio"
	"finance-help/config"
	"finance-help/internal/services"
	"finance-help/internal/web/controllers"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	cfg := config.LoadConfig()
	log.Println("Config Loaded")

	if cfg.ENV == "ngrok" {
		urlChan := make(chan string)
		initNgrok(urlChan)
		setTelegramWebhook(*cfg, urlChan)
	}

	tgService := services.NewTelegramServiceImpl()
	// TODO: temporal, just to simulate whatsapp impl for future reference.
	messageController := controllers.NewMessageServiceController(tgService, tgService)

	http.HandleFunc("/", messageController.HandleEvent)

	log.Printf("Starting server on :%s", cfg.PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.PORT), nil); err != nil {
		log.Fatal("Err starting server: ", err)
	}
}

func initNgrok(urlChan chan string) {
	cmd := exec.Command("ngrok", "http", "8080", "--log=stdout")

	reader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("Error creating reader: ", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting ngrok:", err)
	}

	go func() {
		sc := bufio.NewScanner(reader)
		for sc.Scan() {
			if strings.Contains(sc.Text(), "url=") {
				_, after, _ := strings.Cut(sc.Text(), "url=")
				urlChan <- after
			}
		}
	}()
}

func setTelegramWebhook(cfg config.Config, urlChan chan string) {
	publicURL := <-urlChan

	log.Printf("URL NGROK: %s", publicURL)

	client := &http.Client{}
	url := fmt.Sprintf("%s/bot%s/setWebhook", cfg.TelegramApiEndpoint, cfg.TelegramToken)
	body := fmt.Sprintf(`{"url": "%s"}`, publicURL)

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		log.Fatal("Err creating request: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Err sending request: ", err)
	}

	log.Println("Response status: ", resp.Status)
}
