package config

import {{.authImport}}

var (
    AppConf *Config
)

type Config struct {
	rest.RestConf
	{{.auth}}
	{{.jwtTrans}}
	{{if .mysql}}
	DbConfig DBConfig
	{{end}}
}
