API_REST_OUT := "build-rest/" # TODO - NO API ENDPOINT
PKG := "github.com/olamai/api"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
JS_OUT := "../web-client/lib/proto"
PY_OUT := "../python-client"
.PHONY: all

all: compileGO compileJS

compileGO: # compile proto files
		@protoc \
		-I ./ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:./ \
		./*.proto

compileJS:
		@protoc -I=./ simulation.proto \
		--js_out=import_style=commonjs:${JS_OUT} \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:${JS_OUT}

compilePY:
		@python3 -m grpc_tools.protoc -I./ ./simulation.proto \
		--python_out=${PY_OUT} \
		--grpc_python_out=${PY_OUT}

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