#!/usr/bin/env bash
curl -H "Content-type: application/json" -d '{"firstname":"Bob","lastname":"Jones","address":"123 Any Street","gender":"male"}' http://localhost:8080/person
curl -H "Content-type: application/json" -d '{"firstname":"Joe","lastname":"Smith","address":"3324 Elm Ave","gender":"male","timestamp":"2019-05-03T08:28:27+05:00"}' http://localhost:8080/person

