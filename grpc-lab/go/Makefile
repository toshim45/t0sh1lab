clean:
	rm -rfv messaging
	rm -rfv messagingv1
	rm -rfv binary
proto-build:
	protoc --go_out=. --go-grpc_out=. --proto_path=.. message.proto
	protoc --go_out=. --go-grpc_out=. --proto_path=.. messagev1.proto
build:
	go build -o binary .
runb:
	./binary
