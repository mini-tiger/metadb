{{- define "mytemplate" -}}
{{ println }}
- name: redisIp
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: redisIp
- name: redisDBnum
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: redisDBnum
- name: redisPass
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: redisPass
- name: redisPort
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: redisPort
- name: mongoPort
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: mongoPort
- name: mongoIp
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: mongoIp
- name: mongoRootUser
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: mongoRootUser
- name: mongoRootPass
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: mongoRootPass
- name: mongoRootAuthDB
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: mongoRootAuthDB
- name: mongoShardNode
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: mongoShardNode
- name: mongoCluster
  valueFrom:
    secretKeyRef:
      name: metadb-secret
      key: mongoCluster
{{- end -}}