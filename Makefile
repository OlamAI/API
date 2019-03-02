API_REST_OUT := "build-rest/" # TODO - NO API ENDPOINT
PKG := "github.com/olamai/api"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
JS_OUT := "../web-client/lib/proto"
PY_OUT := "../../python-client/"
.PHONY: all

all: compileGO compileJS compilePY

compileGO: # compile proto files
	@protoc \
	-I ./ \
	-I${GOPATH}/src \
	--include_imports \
	--include_source_info \
	--proto_path=/usr/local/include \
	--proto_path=./simulation \
	--go_out=plugins=grpc:./ \
	--descriptor_set_out=./simulation/api_descriptor.pb	\
	./simulation/simulation.proto

compileJS:
	@protoc -I=./ ./*.proto \
	--proto_path=/usr/local/include \
	--js_out=import_style=commonjs:${JS_OUT} \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:${JS_OUT}

compilePY:
	@python3 -m grpc_tools.protoc -I./ ./*.proto \
		--python_out=. \
		--grpc_python_out=.

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

# python3 -m grpc_tools.protoc -I/usr/local/include/google/api http.proto  --python_out=.
# python3 -m grpc_tools.protoc -I/usr/local/include -I/usr/local/include/google/api annotations.proto  --python_out=.

# python3 -m grpc_tools.protoc -I. http.proto  --python_out=.
# python3 -m grpc_tools.protoc -I. annotations.proto  --python_out=.

protoc ./annotations.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. --descriptor_set_out=.
protoc ./http.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. 
protoc ./descriptor.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. 
protoc ./simulation.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. 