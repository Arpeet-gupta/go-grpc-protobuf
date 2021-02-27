gen:
	protoc -I=proto/ --go_out=. --go_opt=module=github.com/Arpeet-gupta/go-grpc-protobuf proto/*.proto
clean:
	rm pb/*.go

delete:
	rm main
run:
	go build main.go
get:
	protoc --go_out=. --go_opt=paths=source_relative proto/*.proto
