docker-run-prod:
	docker compose -f ./docker-compose-prod.yml up --build -d
docker-run-dev:
	docker compose -f ./docker-compose-dev.yml up --build -d
docker-run-dev-recreate:
	docker compose -f ./docker-compose-dev.yml up --build -d golang