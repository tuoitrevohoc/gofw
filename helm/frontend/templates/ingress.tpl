{{- if .Values.ingress.enabled -}}
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: {{ .Values.fullnameOverride }}
  namespace: {{ .Release.Namespace | default "default" }}
  labels:
    app.kubernetes.io/name: {{ .Values.fullnameOverride }}
    helm.sh/chart: {{ printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
    app.kubernetes.io/instance: {{ .Release.Name }}
    app.kubernetes.io/managed-by: {{ .Release.Service }}
    {{- with .Values.commonLabels }}
    {{- toYaml . | nindent 4 }}
    {{- end }}
  {{- with .Values.ingress.annotations }}
  annotations:
    {{- toYaml . | nindent 4 }}
  {{- end }}
  {{- with .Values.commonAnnotations }}
  {{- toYaml . | nindent 4 }}
  {{- end }}
spec:
  {{- if .Values.ingress.ingressClassName }}
  ingressClassName: {{ .Values.ingress.ingressClassName }}
  {{- end }}
  {{- if .Values.ingress.tls }}
  tls:
    {{- range .Values.ingress.tls }}
    - hosts:
        {{- range .hosts }}
        - {{ . | quote }}
        {{- end }}
      secretName: {{ .secretName }}
    {{- end }}
  {{- end }}
  rules:
    - host: {{ .Values.ingress.hostname | quote }}
      http:
        paths:
          - path: {{ .Values.ingress.path | quote }}
            pathType: {{ .Values.ingress.pathType | default "Prefix" }}
            backend:
              service:
                name: {{ .Values.fullnameOverride }} # This should match your service name
                port:
                  number: {{ .Values.service.port }} # This should match your service port
{{- end }}