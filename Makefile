initialize:
	cd gateway; make initialize
	cd shop; make initialize

protobuf:
	cd gateway; make protobuf
	cd shop; make protobuf