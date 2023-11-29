package {{.pkgName}}

{{if .mysql}}
type DBConfig struct {
	Ip              string
	Port            int
	Pwd             string
	User            string
	ConnectPoolSize int
	SetLog          bool
}
{{end}}

{{if .redis}}
type CacheConfig struct{
    Ip string
    Port int
}
{{end}}