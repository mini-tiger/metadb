### 1. MTDADB

#### 1.1编译镜像

```
cd configcenter/scripts
images.go
 
```

#### 1.2 部署k8s

```
#登录k8s master , clone代码
cd configcenter/src/deploy

helm install mongo -n cmdbv3 -f values.yaml ./mongodb-helm
helm install redis -n cmdbv3 -f values.yaml ./redis-helm
helm install zookeeper -n cmdbv3 -f values.yaml ./zookeeper-helm

添加mongo 用户
kubectl -n cmdbv3 exec mongo-mongodb-primary-0 -ti bash
mongo -u root -p cc
use cmdb
db.createUser({user: "cc",pwd: "cc",roles: [ { role: "readWrite", db: "cmdb" } ]})

cd configcenter/deploy/cmdb-helm
helm install -n cmdbv3 cmdb .
```

#### 1.3 开发环境搭建

./mongod --port 27017 --dbpath /soft/mongodb/data/db0 --replSet rs0 --auth --bind_ip_all --fork --logpath=/data/logs

```shell
cd $GOPATH/src/configcenter/deploy/compose
1.mongo
mongo集群部署

 startdb.sh

mongo 创建用户
 docker exec -ti mongo1 bash
 mongo
 > use cmdb
 > db.createUser({user: "cc",pwd: "cc",roles: [ { role: "readWrite", db: "cmdb" } ]})

mongo 删除
# docker-compose -f mongo-compose.yml down -v
docker stop mongo1
docker stop mongo2
docker stop mongo3
docker rm mongo1
docker rm mongo2
docker rm mongo3
rm -rf /home/taojun/mongors

2. redis
redis 部署
 docker-compose -f redis-compose.yml up -d
 
3. zk
 zk 部署
 docker-compose -f zk-compose.yml up -d
 zk 删除
 # docker-compose -f zk-compose.yml down -v
 rm -rf /home/taojun/zookeeper
```

### 2.ui

#### 2.1 编译镜像

```
cd configcenter/deploy_ui
./docker.sh
```

#### 2.2 部署k8s

```
#登录k8s master , clone代码
cd configcenter/deploy_ui/cmdb-ui
helm install -n cmdb ui
```

#### 2.3 开发环境搭建

```
cd configcenter/indc-front
npm run serve
```