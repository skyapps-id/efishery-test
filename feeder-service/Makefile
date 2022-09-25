run:
	go run cmd/main.go
build:
	go build cmd/main.go
clean:
	go mod tidy
protoc:
	protoc --proto_path=proto ./proto/*.proto --go_out=. --go-grpc_out=.