#!/usr/bin/env bash

docker-compose down
rm -rf data/redis0{1,2,3,4,5,6}/data/*
