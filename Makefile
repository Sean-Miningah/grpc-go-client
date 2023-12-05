gen_grpc_protoc:
	@echo "Generating gRPC protobuf files"
	protoc \
		--proto_path=proto \
		--go_out=protobuffers \
		--go_opt=paths=source_relative \
		--go-grpc_out=protobuffers \
		--go-grpc_opt=paths=source_relative \
		proto/*.proto 


