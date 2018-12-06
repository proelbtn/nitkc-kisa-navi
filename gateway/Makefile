TAG := proelbtn/school-eve-navi-gateway
BARGS := --tag ${TAG}
RARGS := --interactive --tty --publish 8080:80

all:

initialize:
	go get -v -u github.com/golang/protobuf/{proto,protoc-gen-go}

protobuf:
	mkdir -p protos
	protoc --proto_path=../protos --go_out=plugins=grpc:"$(shell go env GOPATH)/src" ../protos/food/service.proto ../protos/food/messages.proto
	protoc --proto_path=../protos --go_out=plugins=grpc:"$(shell go env GOPATH)/src" ../protos/shop/service.proto ../protos/shop/messages.proto
	protoc --proto_path=../protos --go_out=plugins=grpc:"$(shell go env GOPATH)/src" ../protos/souvenir/service.proto ../protos/souvenir/messages.proto

build:
	DOCKER_BUILDKIT=1 docker build ${BARGS} .

run:
	docker run ${RARGS} ${TAG}
