## Compile JS

protoc -I=./ environment.proto \
--js_out=import_style=commonjs:./ \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./

## Compile Python

python3 -m grpc_tools.protoc -I./ --python_out=. --grpc_python_out=. ./environment.proto

# Connect to cluster from kubetcl

gcloud container clusters get-credentials --zone us-central1-a [CLUSTER_NAME]

# Deploy to GKE

kubectl create -f grpc-environment.yaml

# GKE - Generate descriptor

protoc -m grpc_tools.protoc \
 --include_imports \
 --include_source_info \
 --proto_path=. \
 --descriptor_set_out=api_descriptor.pb \
 --python_out=generated_pb2 \
 --grpc_python_out=generated_pb2 \
 bookstore.proto

# Development 
Start redis  
`kubectl run redismod --image=redislabs/redismod --port=6379` 
`kubectl expose deployment redismod --type=NodePort` 
Start environment
`kubectl create -f environment/deployment-staging.yaml` 
`kubectl create -f environment/service.yaml` 
Start esp  
`kubectl create -f esp/deployment-staging.yaml`  
Expose esp
`kubectl create -f esp/service-staging.yaml`
