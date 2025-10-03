# debug time value
time=

# env
include .env
DOCKER_CONTAINER_WEB=nginx
DOCKER_CONTAINER_API=api
DOCKER_CONTAINER_DB=postgres
DB_HOST=localhost

# cmd
up:
	@if [ ! -e ".air.toml" ]; then bash ./docker/api/air.sh ;fi
	docker compose -f docker-compose.yml -f ./docker/docker-compose.${ENV}.yml up -d
down:
	docker compose -f docker-compose.yml -f ./docker/docker-compose.${ENV}.yml down
logs-web:
	docker logs -f ${DOCKER_CONTAINER_WEB}
logs-api:
	docker logs -f ${DOCKER_CONTAINER_API}
logs-db:
	docker logs -f ${DOCKER_CONTAINER_DB}
# e.x.) make db-migrate-create name=migrate-content
db-migrate-create:
	migrate create -ext sql -dir migrations -seq ${name}
db-migrate-up:
	migrate -path migrations -database 'postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=bcrd&sslmode=disable' -verbose up
db-migrate-down:
	migrate -path migrations -database 'postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=bcrd&sslmode=disable' -verbose down
run:
	go run main.go
grpcurl-shops:
	grpcurl -plaintext localhost:58080 bcrd.bcrdService.GetShops
main-branch:
	git switch main
	git branch | grep -v "main" | xargs git branch -d
	git pull
develop-branch:
	git switch develop
	git branch | grep -v "develop" | xargs git branch -d
	git pull
update-gomod:
	go clean -cache -modcache -i -r
	go mod download
	go get -u github.com/go-to/bcrd_protobuf/pb@latest
	go get -u
	go mod tidy
# e.x.) make debug-time time='2025-03-11 19:00:00'
debug-time:
	mysql -u${DB_USER} -p${DB_PASS} -P${DB_PORT} -h${DB_HOST} ${DB_NAME} -e "UPDATE config SET conf_value = '${time}' WHERE conf_name = 'debug_time';"
