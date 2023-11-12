set dotenv-load

# build -> build application
build:
	go build -o main ./cmd

# run -> application
run:
	./main

# dev -> run build then run it
dev: 
	watchexec -r -c -e go -- just build run

# health -> Hit Health Check Endpoint
health:
	curl -s http://localhost:8000/healthz | jq

# migrate-create -> create migration
migrate-create NAME:
	migrate create -ext sql -dir ./migrations -seq {{NAME}}

# migrate-up -> up migration
migrate-up:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/todos?sslmode=disable up

# seed: seed 100 todo
seed:
	k6 run ./scripts/seed.js