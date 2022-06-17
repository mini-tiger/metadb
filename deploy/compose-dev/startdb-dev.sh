#!/bin/bash

docker-compose -f mongo-compose-dev.yml up -d
sleep 5
docker exec mongo1 /scripts/rs-init.sh
sleep 20

docker exec mongo1 /scripts/dbuser-init.sh