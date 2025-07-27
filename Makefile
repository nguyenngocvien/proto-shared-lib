proto_user:
	protoc --proto_path=. --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative --grpc-gateway_out=gen --grpc-gateway_opt=paths=source_relative user/*.proto

proto_group:
	protoc --proto_path=. --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative --grpc-gateway_out=gen --grpc-gateway_opt=paths=source_relative group/*.proto

proto_chat:
	protoc --proto_path=. --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative --grpc-gateway_out=gen --grpc-gateway_opt=paths=source_relative chat/*.proto

proto_workflow:
	protoc --proto_path=. --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative --grpc-gateway_out=gen --grpc-gateway_opt=paths=source_relative workflow/*.proto

proto_swagger:
	protoc --proto_path=. --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=vcore-stack user/*.proto group/*.proto chat/*.proto workflow/*.proto

proto:
	protoc --proto_path=. --go_out=gen --go_opt=paths=source_relative --go-grpc_out=gen --go-grpc_opt=paths=source_relative --grpc-gateway_out=gen --grpc-gateway_opt=paths=source_relative --openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=swagger openapi/*.proto user/*.proto group/*.proto chat/*.proto workflow/*.proto