build:
	docker build -t skill-api:1.0.0 ./api
	docker build -t skill-consumer:1.0.0 ./consumer
run:
	docker compose up -d