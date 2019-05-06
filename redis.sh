#!/usr/bin/env bash

docker run -d --hostname my-redis --name some-redis -p 6379:6379 redis

