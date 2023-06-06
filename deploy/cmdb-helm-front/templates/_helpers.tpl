{{/*
Expand the name of the chart.
*/}}
{{- define "lmanager-cmdb-ui.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "lmanager-cmdb-ui.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := coalesce .Values.nameOverride ((.Values.global).lmanagerCmdb).name .Chart.Name }}
{{- printf $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "lmanager-cmdb-ui.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "lmanager-cmdb-ui.labels" -}}
{{ include "lmanager-cmdb-ui.selectorLabels" . }}
helm.sh/chart: {{ include "lmanager-cmdb-ui.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "lmanager-cmdb-ui.selectorLabels" -}}
app.kubernetes.io/name: {{ include "lmanager-cmdb-ui.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "lmanager-cmdb-ui.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "lmanager-cmdb-ui.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}