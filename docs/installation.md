### 1. metaDB

#### 1.1编译打包镜像

```
cd configcenter
make linux

cd configcenter/scripts
go run images.go 
 
```

#### 1.2 部署k8s
##### 1.2.1 deploy mongodb
```shell
cd configcenter/deploy/mongodb-sharded
helm install -n ${ns} mongodb-sharded
```

##### 1.2.2 deploy redis
```shell
cd configcenter/deploy/redis-helm
helm install -n ${ns} redis
```

#### 1.2.3 deploy metadb
```
#登录k8s master , clone代码
cd configcenter/deploy/cmdb-helm-allinone/cmdb-helm

helm install -n ${ns} cmdb .
```

#### 1.3 开发环境搭建
##### 1.3.1 mongo,redis deploy
基于 docker-compose 3.x
```shell
cd configcenter/deploy/compose-dev
./deploy-dev.sh


```

##### 1.3.2 部署k8s helm metadb
```
#登录k8s master , clone代码
cd configcenter/deploy_ui/cmdb-ui
helm install -n ${ns} cmdb .
```

##### 1.3.3 front ui

```
cd configcenter/indc-front
npm run serve
```