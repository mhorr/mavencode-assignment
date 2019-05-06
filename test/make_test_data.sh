#!/usr/bin/env bash

go get github.com/Pallinder/go-randomdata
go get github.com/gomodule/redigo/redis
go get github.com/spf13/viper
go get github.com/streadway/amqp
go get github.com/mhorr/mavencode-assignment/shared
cd ../shared
go build
#go get github.com/mhorr/mavencode-assignment/shared
cd ../test

go build

./test 1000
