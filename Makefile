all:

initialize:
	cd gateway; make initialize
	cd food; make initialize

protobuf:
	cd gateway; make protobuf
	cd food; make protobuf

build: protobuf
	cd front; make build
	cd gateway; make build
	cd food; make build

start: build
	docker-compose up -d

destroy:
	docker-compose down
