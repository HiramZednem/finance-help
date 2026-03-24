MAIN=cmd/app/main.go
NGROK_DOMAIN=squattily-expensive-danial.ngrok-free.dev

.PHONY: run

default: run

run:
	go run $(MAIN)

ngrok:
	ngrok http --domain=$(NGROK_DOMAIN) 8080