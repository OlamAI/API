# -----
# Compile Descriptors
# -----
compile-descriptor-environment:
	python3 -m grpc_tools.protoc \
    --include_imports \
    --include_source_info \
    --proto_path=. \
    --proto_path=./environment \
    --descriptor_set_out=./environment/descriptor.pb \
    environment.proto

compile-descriptors: compile-descriptor-environment

# -----
# Compile Go
# -----
compile-go-environment:
	protoc \
    -I=./environment \
	--go_out=plugins=grpc:../simulation/pkg/api/environment \
	environment.proto \

compile-go: compile-go-environment

# -----
# Compile JS
# -----
compile-js-environment:
	cp ./environment/environment.proto ../nodeExample/api
# -----
# Compile
# -----
compile: compile-descriptors compile-go 

# -----
# Deploy
# -----

deploy-environment:
	gcloud endpoints services deploy ./environment/descriptor.pb ./environment/config.yaml

deploy: deploy-environment