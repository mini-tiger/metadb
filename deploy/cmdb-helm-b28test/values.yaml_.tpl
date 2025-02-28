apps:
  - name: adminserver
    port: 60004
#    nodePort: 30971
  - name: webserver
    ports:
      - name: http
        port: 8090
#        nodePort: 32168
      - name: extend
        port: 8081
#        nodePort: 32165
  - name: apiserver
    port: 8080
#    nodePort: 31922
  - name: coreservice
    port: 50009
#    nodePort: 32006
    cacheAffinity: false
  - name: toposerver
    port: 60002
  - name: hostserver
    port: 60001
  - name: operationserver
    port: 60011
  - name: cacheservice
    port: 50010
  - name: cloudserver
    port: 60013
  - name: eventserver
    port: 60009
  - name: procserver
    port: 60003
  - name: taskserver
    port: 60012

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
  redisDBnum: "7"

  #mongo public params
  mongorootuser: "root"
  mongorootauthdb: "admin"
  mongorootpass: "cc"

  # cluster shard

  #mongo_shard_node: "172.60.3.144:27000,172.60.3.145:27000,172.60.3.146:27000"   # b28-test
  mongo_shard_node: "172.22.50.25:32085,172.22.50.25:32086,172.22.50.25:32087"   # 50.25 cmdb-neolink
  #mongo_shard_node: "172.22.50.25:32082,172.22.50.25:32083,172.22.50.25:32084"
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
