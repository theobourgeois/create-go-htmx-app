.PHONY: run run-gen db-u clean

db-up:
	$(eval DB_NAME=$(shell grep DB_NAME .env | cut -d '=' -f2))
	psql -d $(DB_NAME) -f database/schema.sql

run-gen:
	npx tailwindcss -i ./styles/input.css -o ./styles/output.css
	templ generate
	go run internal/main/main.go

run: 
	go run internal/main/main.go

clean: 
	go clean -cache