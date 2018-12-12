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
	cd mysql; make build

start: build
	docker-compose up -d

destroy:
	docker-compose down
