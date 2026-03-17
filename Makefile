MAIN=cmd/app/main.go

.PHONY: run

default: run

run:
	go run $(MAIN)

ngrok:
	ngrok http 8080