TAG := proelbtn/school-eve-navi-food
BARGS := --tag ${TAG}
RARGS := --interactive --tty --publish 30001:30001

all:

initialize:
	pip install grpcio-tools

protobuf:
	python -m grpc_tools.protoc --proto_path=../protos --python_out=. --grpc_python_out=. ../protos/food/messages.proto ../protos/food/service.proto

build:
	DOCKER_BUILDKIT=1 docker build ${BARGS} .

run:
	docker run ${RARGS} ${TAG}
