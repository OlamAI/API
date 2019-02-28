## Compile JS

protoc -I=./ simulation.proto \
--js_out=import_style=commonjs:./ \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./

## Compile Python

python3 -m grpc_tools.protoc -I./ --python_out=. --grpc_python_out=. ./simulation.proto

# Connect to cluster from kubetcl

gcloud container clusters get-credentials --zone us-central1-a [CLUSTER_NAME]

# GKE - Generate descriptor

protoc -m grpc_tools.protoc \
 --include_imports \
 --include_source_info \
 --proto_path=. \
 --descriptor_set_out=api_descriptor.pb \
 --python_out=generated_pb2 \
 --grpc_python_out=generated_pb2 \
 bookstore.proto
