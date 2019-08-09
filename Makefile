JS_OUT = ../web-client/api
PY_OUT = ../pip-pkg/terrariumai
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
# Compile Python
# -----
compile-py-collective:
	python3 -m grpc_tools.protoc \
    -I./collective \
    --python_out=${PY_OUT} \
    --grpc_python_out=${PY_OUT} \
    collective.proto

# -----
# Compile JS
# -----
compile-js-environment:
	protoc \
    -I=./environment \
    --js_out=import_style=commonjs:${JS_OUT} \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:${JS_OUT} \
	environment.proto 

compile-js: compile-js-environment
# -----
# Compile
# -----
compile: compile-descriptors compile-go compile-js

# -----
# Deploy
# -----

deploy-environment-staging:
	gcloud config set project olamai-testing
	gcloud endpoints services deploy ./environment/descriptor.pb ./environment/config-staging.yaml

deploy-environment-prod:
	gcloud config set project olamai-d64a7
	gcloud endpoints services deploy ./environment/descriptor.pb ./environment/config.yaml