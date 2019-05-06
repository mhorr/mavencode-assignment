#!/usr/bin/env bash
curl -H "Content-type: application/json" -d '{"firstname":"Bob","lastname":"Jones","address":"123 Any Street","gender":"male"}' http://localhost:8080/person
