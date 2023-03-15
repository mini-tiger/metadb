apps:
  - name: adminserver
    port: 60004
#    nodePort: 30971
    enable: false
  - name: webserver
    port: 8090
    nodePort: 32168
    enable: true
  - name: apiserver
    port: 8080
    enable: true
    nodePort: 31922
  - name: coreservice
    port: 50009
#    nodePort: 32006
    cacheAffinity: false
    enable: false
  - name: toposerver
    port: 60002
    enable: false
  - name: hostserver
    port: 60001
    enable: false
  - name: operationserver
    port: 60011
    enable: false
  - name: cacheservice
    port: 50010
    enable: false
  - name: cloudserver
    port: 60013
    enable: false
  - name: eventserver
    port: 60009
    enable: false
  - name: procserver
    port: 60003
    enable: false
  - name: taskserver
    port: 60012
    enable: false

prefix: lmanager-cmdb
ingress:
    enable: false

image:
  repository: harbor.dev.21vianet.com/cmdb/
  tag: {{.version}}
#  tag: "latest"
env:
#  pullPolicy: IfNotPresent
  pullPolicy: Always
  redisip: "redis-master"
  #redisip: "172.60.3.120"   # b28-test
  #redisip: "redis-master"   # 50.25 cmdb-neolink
  redispass: "Ne01ink2022!"
  redisport: "6379"
  redisDBnum: "0"

  #mongo public params
  mongorootuser: "root"
  mongorootauthdb: "admin"
  mongorootpass: "cc"

  # cluster shard

  #mongo_shard_node: "172.60.3.144:27000,172.60.3.145:27000,172.60.3.146:27000"   # b28-test
  #mongo_shard_node: "172.22.50.25:32085,172.22.50.25:32086,172.22.50.25:32087"   # 50.25 cmdb-neolink
  mongo_shard_node: "mongo-shard-mongodb-sharded-0:27017,mongo-shard-mongodb-sharded-1:27017,mongo-shard-mongodb-sharded-2:27017"
  #mongo_shard_node: "172.22.50.25:32082,172.22.50.25:32083,172.22.50.25:32084"   #cmdbv4
  mongo_cluster: "shard"
  # 副本集
  mongoip: "mongo-mongodb-headless"
  mongoport: "27017"

# cache coresvc 同一node提高效率
cache:
  enabled: false
  labels:
    - name: app.kubernetes.io/name
      value: redis
    - name: app.kubernetes.io/component
      value: master
#res:
#  req:
#    mem: 1024Mi
#    cpu: 600m
#  req2:
#    mem: 2048Mi
#  req3:
#    mem: 3072Mi
#  ave:
#    mem: 1024Mi
#    cpu: 1000m
