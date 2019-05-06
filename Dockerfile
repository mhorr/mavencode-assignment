FROM golang:latest
ADD . /go/src/github.com/mhorr/mavencode-assignment
RUN go get github.com/gomodule/redigo/redis
RUN go get github.com/streadway/amqp
RUN go get github.com/gorilla/mux
RUN go get github.com/Pallinder/go-randomdata
RUN go get github.com/spf13/viper
WORKDIR /go/src/github.com/mhorr/mavencode-assignment/shared
RUN go build 
WORKDIR /go/src/github.com/mhorr/mavencode-assignment/redisclient
RUN go build 
WORKDIR /go/src/github.com/mhorr/mavencode-assignment/webserver
RUN go build
WORKDIR /go/src/github.com/mhorr/mavencode-assignment/test
RUN go build
WORKDIR /go/src/github.com/mhorr/mavencode-assignment