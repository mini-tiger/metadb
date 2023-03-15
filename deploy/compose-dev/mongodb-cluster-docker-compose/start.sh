#!/bin/bash

docker-compose up -d
sleep 5
docker-compose exec configsvr01 sh -c "mongosh < /scripts/init-configserver.js"
sleep 1
docker-compose exec shard01-a sh -c "mongosh < /scripts/init-shard01.js"
sleep 1
docker-compose exec shard02-a sh -c "mongosh < /scripts/init-shard02.js"
sleep 1
docker-compose exec shard03-a sh -c "mongosh < /scripts/init-shard03.js"
sleep 1
docker-compose exec router01 sh -c "mongosh < /scripts/init-runcommand.js"
sleep 1
docker-compose exec router01 sh -c "mongosh < /scripts/init-router.js"
sleep 1
docker-compose exec router01 sh -c "mongosh < /scripts/init-user.js"

sleep 1

docker exec -it mongo-config-01 bash -c "echo 'rs.status()' | mongosh --port 27017"