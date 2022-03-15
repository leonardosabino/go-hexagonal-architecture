all: build run

build:
	go build ./cmd

run:
	go run ./cmd

clean:
	go mod tidy

test:
	go test -v -count=1 ./...

test-cover:
	go test -v -count=1 ./... -covermode=count  -json | tparse -top

generate-mocks:
	cd ./internal && mockery --all --keeptree


##### DOCKER FILE #####

# load pg changes in sql folder
rebuild-env: stop-env destroy-env build-env

start-env:
	docker-compose -f ./build/package/docker/localstack/docker-compose-local.yml up

stop-env:
	docker-compose -f ./build/package/docker/localstack/docker-compose-local.yml stop

destroy-env:
	docker-compose -f ./build/package/docker/localstack/docker-compose-local.yml rm -f

build-env:
	docker-compose -f ./build/package/docker/localstack/docker-compose-local.yml build