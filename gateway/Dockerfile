FROM golang:alpine

RUN apk add git

ADD . /go/src/github.com/proelbtn/school-eve-navi/gateway
WORKDIR /go/src/github.com/proelbtn/school-eve-navi/gateway

RUN go get -v .

CMD [ "go", "run", "-v", "main.go" ]