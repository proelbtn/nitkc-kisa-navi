all:

initialize:
	cd gateway; make initialize
	cd shop; make initialize

protobuf:
	cd gateway; make protobuf
	cd shop; make protobuf

build: protobuf
	cd gateway; make build
	cd shop; make build

run: build
	docker-compose up
	docker-compose down