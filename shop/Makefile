all: run

initialize:
	pip install grpcio-tools

protobuf:
	python -m grpc_tools.protoc --proto_path=../protos --python_out=. --grpc_python_out=. ../protos/food/messages.proto ../protos/food/service.proto

run:
	python main.py
