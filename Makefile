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
compile-descriptor-collective:
	python3 -m grpc_tools.protoc \
    --include_imports \
    --include_source_info \
    --proto_path=. \
    --proto_path=./collective \
    --descriptor_set_out=./collective/descriptor.pb \
    collective.proto

compile-descriptors: compile-descriptor-collective compile-descriptor-environment

# -----
# Compile Go
# -----
compile-go-environment:
	protoc \
    -I=./environment \
	--go_out=plugins=grpc:../simulation/pkg/api/environment \
	environment.proto 
compile-go-collective:
	protoc \
    -I=./collective \
	--go_out=plugins=grpc:../simulation/pkg/api/collective \
	collective.proto \

compile-go: compile-go-collective compile-go-environment

# -----
# Compile JS
# -----
# Note: JS can read in .proto files directly so we smiply copy it over
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

deploy-collective:
	gcloud endpoints services deploy ./environment/descriptor.pb ./environment/config.yaml

deploy: deploy-environment