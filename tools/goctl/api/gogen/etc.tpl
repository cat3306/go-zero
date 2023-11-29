Name: {{.serviceName}}
Host: {{.host}}
Port: {{.port}}
{{if .mysql}}
DbConfig:
  Ip: 127.0.0.1
  Port: 3306
  Pwd: "12345678"
  User: root
  ConnectPoolSize: 100
  SetLog: true
{{end}}