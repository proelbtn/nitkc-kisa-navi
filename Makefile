all:

initialize:
	cd gateway; make initialize
	cd food; make initialize

protobuf:
	cd gateway; make protobuf
	cd food; make protobuf

build: protobuf
	cd gateway; make build
	cd food; make build

run: build
	docker-compose up
	docker-compose down