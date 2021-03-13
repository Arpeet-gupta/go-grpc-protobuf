gen:
	protoc -I=proto/ \
	--go_out=. --go_opt=module=github.com/Arpeet-gupta/go-grpc-protobuf \
	--go-grpc_out=. --go-grpc_opt=module=github.com/Arpeet-gupta/go-grpc-protobuf \
	--grpc-gateway_out=. --grpc-gateway_opt=module=github.com/Arpeet-gupta/go-grpc-protobuf \
	--openapiv2_out=swagger \
	proto/*.proto proto/google/api/*.proto
clean:
	rm pb/*.go
delete:
	rm main
server:
	go run cmd/server/main.go --port 8080 

rest:
	go run cmd/server/main.go --port 8082 --type rest --endpoint 192.168.0.121:8080

client:
	go run cmd/client/main.go --address 0.0.0.0:8080

test-all:
	# ./... matches all the packages in the module
	gotest -cover ./...
	# gotest -cover github.com/Arpeet-gupta/go-grpc-protobuf/v6/service/ github.com/Arpeet-gupta/go-grpc-protobuf/v6/serializer
test-protobuf-serializer:
	gotest -cover github.com/Arpeet-gupta/go-grpc-protobuf/v6/serializer/
test-grpc:
	gotest -cover github.com/Arpeet-gupta/go-grpc-protobuf/v6/service/
test-grpc-server:
	gotest -cover -run TestServerCreateLaptop github.com/Arpeet-gupta/go-grpc-protobuf/v6/service/
test-grpc-client:
	# gotest -cover -run TestClientCreateLaptop github.com/Arpeet-gupta/go-grpc-protobuf/v6/service/
	gotest -cover -run TestClientUploadImage github.com/Arpeet-gupta/go-grpc-protobuf/v6/service/