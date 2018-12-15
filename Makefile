all:

initialize:
	cd front; make initialize
	cd gateway; make initialize
	cd food; make initialize
	cd shop; make initialize
	cd souvenir; make initialize

protobuf:
	cd gateway; make protobuf
	cd food; make protobuf
	cd shop; make protobuf
	cd souvenir; make protobuf

build: protobuf
	cd front; make build
	cd gateway; make build
	cd food; make build
	cd shop; make build
	cd souvenir; make build

up:
	docker-compose up

down:
	docker-compose down
