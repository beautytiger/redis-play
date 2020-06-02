#!/usr/bin/env bash

docker-compose down
rm -rf data/redis01/data/* data/redis02/data/* data/redis03/data/*
