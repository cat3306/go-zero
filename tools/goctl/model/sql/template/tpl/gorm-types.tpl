
const (
    {{.upperStartCamelObject}}TName = "{{.tableName}}"
)

type (
	{{.upperStartCamelObject}} struct {
		{{.fields}}
	}
)