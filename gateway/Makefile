all: fmt run

ensure:
	dep ensure

genpb:
	mkdir -p protos/{food,shop,souvenir}
	protoc --proto_path=../protos --go_out=plugins=grpc:protos/food ../protos/food.proto


fmt:
	go fmt ./...

run: ensure
	go run main.go


