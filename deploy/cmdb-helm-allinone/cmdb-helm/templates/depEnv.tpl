{{- define "mytemplate" -}}
{{ println }}
- name: redisIp
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: redisIp
- name: redisDBnum
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: redisDBnum
- name: redisPass
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: redisPass
- name: redisPort
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: redisPort
- name: mongoPort
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: mongoPort
- name: mongoIp
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: mongoIp
- name: mongoRootUser
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: mongoRootUser
- name: mongoRootPass
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: mongoRootPass
- name: mongoRootAuthDB
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: mongoRootAuthDB
- name: mongoShardNode
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: mongoShardNode
- name: mongoCluster
  valueFrom:
    secretKeyRef:
      name: {{ include "lmanager-cmdb.fullname" . }}-secret
      key: mongoCluster
{{- end -}}