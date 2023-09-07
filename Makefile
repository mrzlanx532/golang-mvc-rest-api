docker-prod-up:
	docker compose -f ./docker-compose-prod.yml up --build -d
docker-prod-down:
	docker compose -f ./docker-compose-prod.yml down
docker-dev-up:
	docker compose -f ./docker-compose-dev.yml up --build -d
docker-dev-down:
	docker compose -f ./docker-compose-dev.yml down
docker-dev-recreate:
	docker compose -f ./docker-compose-dev.yml up --build golang