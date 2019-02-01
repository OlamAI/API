## Compile JS

protoc -I=./ simulation.proto \
--js_out=import_style=commonjs:./ \
--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./
