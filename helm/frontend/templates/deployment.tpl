apiVersion: apps/v1
kind: Deployment
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
spec:
  replicas: {{ .Values.replicaCount }}
  selector:
    matchLabels:
      app.kubernetes.io/name: {{ .Values.fullnameOverride | default .Chart.Name }}
      app.kubernetes.io/instance: {{ .Release.Name }}
  template:
    metadata:
      annotations:
        {{- with .Values.commonAnnotations }}
        {{- toYaml . | nindent 8 }}
        {{- end }}
      labels:
        app.kubernetes.io/name: {{ .Values.fullnameOverride | default .Chart.Name }}
        app.kubernetes.io/instance: {{ .Release.Name }}
        {{- with .Values.commonLabels }}
        {{- toYaml . | nindent 8 }}
        {{- end }}
    spec:
      {{- if .Values.serviceAccount.create }}
      serviceAccountName: {{ .Values.serviceAccount.name }}
      {{- end }}
      {{- if .Values.podSecurityContext.enabled }}
      securityContext:
        {{- $psc := deepCopy .Values.podSecurityContext }}
        {{- $_ := unset $psc "enabled" }}
        {{- toYaml $psc | nindent 8 }}
      {{- end }}      
      containers:
        - name: {{ .Chart.Name }}
          image: "{{- if .Values.image.registry }}{{ .Values.image.registry }}/{{ end }}{{ .Values.image.repository }}:{{ .Values.image.tag }}"
          imagePullPolicy: {{ .Values.image.pullPolicy }}
          {{- if .Values.containerSecurityContext }}
          securityContext:
            {{- toYaml .Values.containerSecurityContext | nindent 12 }}
          {{- end }}
          ports:
            {{- if .Values.containerPorts }}
            {{- range $name, $port := .Values.containerPorts }}
            - name: {{ $name }}
              containerPort: {{ $port }}
              protocol: TCP
            {{- end }}
            {{- end }}
          {{- if .Values.livenessProbe.enabled }}
          livenessProbe:
            httpGet:
              path: {{ .Values.livenessProbe.path }}
              port: {{ .Values.livenessProbe.port }}
            initialDelaySeconds: {{ .Values.livenessProbe.initialDelaySeconds | default 15 }}
            periodSeconds: {{ .Values.livenessProbe.periodSeconds | default 10 }}
            timeoutSeconds: {{ .Values.livenessProbe.timeoutSeconds | default 5 }}
            successThreshold: {{ .Values.livenessProbe.successThreshold | default 1 }}
            failureThreshold: {{ .Values.livenessProbe.failureThreshold | default 3 }}
          {{- end }}
          {{- if .Values.readinessProbe.enabled }}
          readinessProbe:
            httpGet:
              path: {{ .Values.readinessProbe.path }}
              port: {{ .Values.readinessProbe.port }}
            initialDelaySeconds: {{ .Values.readinessProbe.initialDelaySeconds | default 5 }}
            periodSeconds: {{ .Values.readinessProbe.periodSeconds | default 10 }}
            timeoutSeconds: {{ .Values.readinessProbe.timeoutSeconds | default 5 }}
            successThreshold: {{ .Values.readinessProbe.successThreshold | default 1 }}
            failureThreshold: {{ .Values.readinessProbe.failureThreshold | default 3 }}
          {{- end }}
          resources:
            {{- toYaml .Values.resources | nindent 12 }}
          volumeMounts:
            - name: reactjs-app-nginx-default-config-volume
              mountPath: /etc/nginx/conf.d
      {{- with .Values.nodeSelector }}
      nodeSelector:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.affinity }}
      affinity:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      {{- with .Values.tolerations }}
      tolerations:
        {{- toYaml . | nindent 8 }}
      {{- end }}
      volumes: 
        - name: reactjs-app-nginx-default-config-volume
          configMap:
            name: reactjs-app-nginx-default-conf
            items:
              - key: default.conf
                path: default.conf
              - key: health-check.conf
                path: health-check.conf