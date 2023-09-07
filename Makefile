docker-run-prod:
	docker compose -f ./docker-compose-prod.yml up --build
docker-run-dev:
	docker compose -f ./docker-compose-dev.yml up --build