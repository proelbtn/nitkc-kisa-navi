all: fmt run

initialize:
	go get -v -u github.com/golang/protobuf/{proto,protoc-gen-go}

protobuf:
	mkdir -p protos
	protoc --proto_path=../protos --go_out=plugins=grpc:protos ../protos/food/service.proto ../protos/food/messages.proto

run:
	go run main.go


