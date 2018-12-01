all: fmt run

initialize:
	go get -v -u github.com/golang/protobuf/{proto,protoc-gen-go}

protobuf:
	mkdir -p protos/{food,shop,souvenir}
	protoc --proto_path=../protos --go_out=plugins=grpc:protos/food ../protos/food.proto

run:
	go run main.go


