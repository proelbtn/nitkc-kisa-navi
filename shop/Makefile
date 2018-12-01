all: run

init:
	pip install grpcio-tools googleapis-common-protos

genpb:
	python -m grpc_tools.protoc --proto_path=../protos --python_out=. --grpc_python_out=. ../protos/food.proto

run:
	python main.py
