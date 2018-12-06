TAG := proelbtn/school-eve-navi-souvenir
BARGS := --tag ${TAG}
RARGS := --interactive --tty --publish 9003:80

all:

initialize:
	pip install grpcio-tools

protobuf:
	python -m grpc_tools.protoc --proto_path=../protos --python_out=. --grpc_python_out=. \
		../protos/souvenir/messages.proto \
		../protos/souvenir/service.proto

build:
	DOCKER_BUILDKIT=1 docker build ${BARGS} .

run:
	docker run ${RARGS} ${TAG}
