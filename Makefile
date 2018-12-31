API_REST_OUT := "build-rest/" # TODO - NO API ENDPOINT
PKG := "github.com/olamai/api"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all

all: compile

compile: # compile proto files
		@protoc \
		-I ./ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:./ \
		./*.proto

dep: ## Get the dependencies
	@go get -v -d ./...

clean: ## Remove previous builds
	@rm ./*.pb.go

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

## ----------------------
## ------ DOCKER --------
## ----------------------
docker-build: ## Build the docker container
	docker build -f Dockerfile.dev -t api .

docker-up: default ## Start the service locally
	@echo "=============starting api locally============="
	docker-compose up -d

docker-logs: ## tail the logs
	docker-compose logs -f

docker-down: ## shut down the docker server
	docker-compose down

docker-clean: down ## shuts down the API and then clears out saved Docker images 
	@echo "=============cleaning up============="
	rm -f api
	docker system prune -f
	docker volume prune -f