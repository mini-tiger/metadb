

nameOverride: ""
fullnameOverride: ""


serviceAccount:
  # Specifies whether a service account should be created
  create: false
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""

autoScaling:
  enabled: false

# 生产修改实例数
replicaCount: 1

ingress:
    enable: false

image:
  repository: {{.harborDomain}}/lclouds/cmdb_allinone
  tag: {{.version}}
  pullPolicy: Always


initContainers:
  #mongoimage: {{.harborDomain}}/lclouds/mongo:v4.4.13-debian-10-r52
  redisimage: {{.harborDomain}}/lclouds/redis:vbitnami-6.2.6-debian-10-r103
  jobmongoimage: {{.harborDomain}}/lclouds/mongodb-sharded:v6.0.1-debian-11-r10
  jobcurl: {{.harborDomain}}/lclouds/curl:vcentos

imagePullSecrets: []

env:
  pullPolicy: Always
  redisuser: "neolink"
{{ if or (contains .env "cmdbv4") (contains .env "neolink") }}
  redisip: "redis-master"
{{ end }}
{{ if (contains .env "b28-test") }}
  redisip: "172.60.3.120"
{{ end }}
{{ if (contains .env "m6") }}
  redisip: "172.22.34.22"
{{ end }}

{{ if (contains .env "m6") }}
  redispass: "ne0liNk2022#"
  redisport: "6379"
  redisDBnum: "7"
{{ else }}
  redispass: "Ne01ink2022!"
  redisport: "6379"
  redisDBnum: "2"
{{ end }}

  #mongo public params
{{ if (contains .env "m6") }}
  mongorootuser: "neolink"
  mongorootpass: "Ne01ink2022!"
{{ end }}


{{ if (contains .env "b28-test") }}
  mongorootuser: "neolink"
  mongorootpass: "Ne01ink2022!"
{{ end }}

{{ if or (contains .env "cmdbv4") (contains .env "neolink") }}
  mongorootuser: "root"
  mongorootpass: "cc"
{{ end }}
  mongorootauthdb: "admin"


  # cluster shard
{{ if (contains .env "cmdbv4") }}
  mongo_shard_node: "172.22.50.25:32082,172.22.50.25:32083,172.22.50.25:32084"    # cmdbv4
{{ end }}
{{ if (contains .env "neolink") }}
  mongo_shard_node: "172.22.50.25:32085,172.22.50.25:32086,172.22.50.25:32087"   # 50.25 cmdb-neolink
{{ end }}
{{ if (contains .env "b28-test") }}
  mongo_shard_node: "172.60.3.144:27000,172.60.3.145:27000,172.60.3.146:27000"   # b28-test
{{ end }}
{{ if (contains .env "m6") }}
  mongo_shard_node: "172.22.34.31:27000,172.22.34.32:27000,172.22.34.33:27000"   # m6
{{ end }}
  mongo_cluster: "shard"
  # 副本集
  mongoip: "mongo-mongodb-headless"
  mongoport: "27017"

podAnnotations: {}

podSecurityContext: {}
# fsGroup: 2000

securityContext: {}
  # capabilities:
  #   drop:
  #   - ALL
  # readOnlyRootFilesystem: true
  # runAsNonRoot: true
# runAsUser: 1000

service:
  - type: NodePort
    port: 8090
    targetPort: 8090
    nodePort: 32178
    name: webserver
  - type: NodePort
    port: 8080
    targetPort: 8080
    name: apiserver
{{ if (contains .env "m6") }}
    nodePort: 31882   # m6
{{ else }}
    nodePort: 31882
{{ end }}

adminserver:
    type: ClusterIP
    port: 60004
    targetPort: 60004
    nodePort: 0
    name: adminserver

containerPorts:
  webserver: 8090
  apiserver: 8080



# 存活探针
livenessProbe:
  enabled: true
  initialDelaySeconds: 120
  periodSeconds: 10
  timeoutSeconds: 5
  successThreshold: 1
  failureThreshold: 3
# 就绪探针
readinessProbe:
  enabled: true
  initialDelaySeconds: 60
  periodSeconds: 10
  timeoutSeconds: 1
  successThreshold: 1
  failureThreshold: 3

# 资源限制
resources:
  limits:
    cpu: 4
    memory: 4Gi
  requests:
    cpu: 2
    memory: 4Gi

# 标签选择器
nodeSelector:
  kubernetes.io/os: linux

# 资源调度
tolerations: []

global:
  ui:
    enabled: false
    replicaCount: 1
    repository: harbor.dev.21vianet.com/lclouds/cmdb_matedb_ui
    tag: v1.1.0
    service:
      type: NodePort
      port: 80
      targetPort: 80
      nodePort: 32174
    ingress:
      enabled: false
      hosts:
        - paths:
            - path: /lmanager-cmdb-ui
              pathType: Prefix
    resources:
      limits:
        cpu: 2
        memory: 2Gi
      requests:
        cpu: 0.5
        memory: 1Gi

