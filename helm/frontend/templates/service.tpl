{{- if .Values.service.enabled -}}
apiVersion: v1
kind: Service
metadata:
  name: {{ .Values.fullnameOverride | default .Chart.Name }}
  namespace: {{ .Release.Namespace }}
  labels:
    app.kubernetes.io/name: {{ .Values.fullnameOverride | default .Chart.Name }}
    helm.sh/chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+" "_" }}
    app.kubernetes.io/instance: {{ .Release.Name }}
    app.kubernetes.io/managed-by: {{ .Release.Service }}
    {{- with .Values.commonLabels }}
    {{- toYaml . | nindent 4 }}
    {{- end }}
  {{- with .Values.service.annotations }}
  annotations:
    {{- toYaml . | nindent 4 }}
  {{- end }}
spec:
  type: {{ .Values.service.type }}
  ports:
    - port: {{ .Values.service.port }}
      targetPort: {{ .Values.service.targetPort | default "http" }} # Default to named port 'http' from deployment
      protocol: TCP
      name: http
  selector:
    app.kubernetes.io/name: {{ .Values.fullnameOverride | default .Chart.Name }}
    app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}