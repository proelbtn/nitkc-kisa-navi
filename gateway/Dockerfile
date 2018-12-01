FROM golang

ADD . /go/src/github.com/proelbtn/school-eve-navi/gateway
RUN go get -v /go/src/github.com/proelbtn/school-eve-navi/gateway

WORKDIR /go/src/github.com/proelbtn/school-eve-navi/gateway

CMD [ "go", "run", "-v", "main.go" ]