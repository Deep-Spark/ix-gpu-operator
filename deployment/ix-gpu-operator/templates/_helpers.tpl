{{/*
Expand the name of the chart.
*/}}
{{- define "ix-gpu-operator.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "ix-gpu-operator.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "ix-gpu-operator.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "ix-gpu-operator.labels" -}}
helm.sh/chart: {{ include "ix-gpu-operator.chart" . }}
{{ include "ix-gpu-operator.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "ix-gpu-operator.selectorLabels" -}}
app.kubernetes.io/name: {{ include "ix-gpu-operator.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "ix-gpu-operator.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "ix-gpu-operator.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{- define "ix-gpu-operator.operand-labels" -}}
helm.sh/chart: {{ include "ix-gpu-operator.chart" . }}
app.kubernetes.io/managed-by: {{ include "ix-gpu-operator.name" . }}
{{- if .Values.daemonsets.labels }}
{{ toYaml .Values.daemonsets.labels | nindent 4 }}
{{- end }}
{{- end }}

{{/*
Full image name with tag
*/}}
{{- define "ix-gpu-operator.image" -}}
{{- .Values.operator.image -}}:{{- .Values.operator.tag | default .Chart.AppVersion -}}
{{- end }}