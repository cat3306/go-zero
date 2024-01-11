package gen

import (
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func genTag(table Table, in string) (string, error) {
	if in == "" {
		return in, nil
	}

	text, err := pathx.LoadTemplate(category, tagTemplateFile, template.Tag)
	if err != nil {
		return "", err
	}

	output, err := util.With("tag").Parse(text).Execute(map[string]any{
		"field": in,
		"data":  table,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func genGormTag(table Table, field *parser.Field) (string, error) {
	in := field.NameOriginal
	if in == "" {
		return in, nil
	}
	text, err := pathx.LoadTemplate(category, tagGormTemplateFile, template.GormTag)
	if err != nil {
		return "", err
	}

	output, err := util.With("tag").Parse(text).Execute(map[string]any{
		"field":                    in,
		"isTimeAndHasDefaultValue": field.IsTime && field.HasDefaultValue,
		"data":                     table,
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
