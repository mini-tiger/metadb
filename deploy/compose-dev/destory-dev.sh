#!/bin/bash
docker-compose -f mongo-compose-dev.yml down
rm -rf /home/taojun/mongors

#docker stop mongo1
#docker stop mongo2
#docker stop mongo3
#docker rm mongo1
#docker rm mongo2
#docker rm mongo3

docker-compose -f redis-compose.yml down
#docker-compose -f meta-svc.yml down -v --rmi all
docker network rm work-net-dev