package config

import {{.authImport}}

var (
    AppConf *Config
)
type HandlerConf struct {
	HandlerErrLog bool
}
type Config struct {
	rest.RestConf
	HandlerConf

	{{.auth}}
	{{.jwtTrans}}
	{{if .mysql}}
	DbConfig DBConfig
	{{end}}

	{{if .redis}}
    Redis RedisConfig
    {{end}}
}
