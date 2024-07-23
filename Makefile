build:
	docker build -t skill-api:1.0.0 ./api
	docker build -t skill-consumer:1.0.0 ./consumer
run:
	docker compose up -d
push:
	docker tag skill-api:1.0.0 ghcr.io/putitat/skill-api:latest
	docker push ghcr.io/putitat/skill-api:latest

	docker tag skill-consumer:1.0.0 ghcr.io/putitat/skill-consumer:latest
	docker push ghcr.io/putitat/skill-consumer:latest