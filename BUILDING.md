Building the project:
go get github.com/gorilla/mux
go get github.com/streadway/amqp
go get github.com/gomodule/redigo/redis
cd src/redisclient
go build
cd ../webserver
go build
start rabbit container (rabbit.sh)
start redis container (redis.sh)
start webserver
start redisclient
