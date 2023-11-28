package gen

import (
	"bytes"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/model"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"github.com/zeromicro/go-zero/tools/goctl/util/stringx"
	"os"
	"path/filepath"
)

type GormGenerator struct {
	generatorConf
}
type gormCode struct {
	importsCode string
	typesCode   string
	methodCode  string
}

// NewGormGenerator creates an instance for defaultGenerator
func NewGormGenerator(dir string, cfg *config.Config, opt ...Option) (*GormGenerator, error) {
	if dir == "" {
		dir = pwd
	}
	dirAbs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	dir = dirAbs
	pkg := util.SafeString(filepath.Base(dirAbs))
	err = pathx.MkdirIfNotExist(dir)
	if err != nil {
		return nil, err
	}
	generator := &GormGenerator{generatorConf{
		dir: dir, cfg: cfg, pkg: pkg,
	}}
	var optionList []Option
	optionList = append(optionList, newDefaultOption())
	optionList = append(optionList, opt...)
	for _, fn := range optionList {
		fn(&generator.generatorConf)
	}

	return generator, nil
}
func (g *GormGenerator) genFromDDL(filename string, withCache, strict bool, database string) (
	map[string]*codeTuple, error,
) {
	m := make(map[string]*codeTuple)
	tables, err := parser.Parse(filename, database, strict)
	if err != nil {
		return nil, err
	}

	for _, e := range tables {
		code, err := g.genModel(*e, withCache)
		if err != nil {
			return nil, err
		}
		//customCode, err := g.genModelCustom(*e, withCache)
		//if err != nil {
		//	return nil, err
		//}

		m[e.Name.Source()] = &codeTuple{
			modelCode: code,
		}
	}

	return m, nil
}
func (g *GormGenerator) StartFromInformationSchema(tables map[string]*model.Table, withCache, strict bool) error {
	m := make(map[string]*codeTuple)
	for _, each := range tables {
		table, err := parser.ConvertDataType(each, strict)
		if err != nil {
			return err
		}

		code, err := g.genModel(*table, withCache)
		if err != nil {
			return err
		}
		customCode, err := g.genModelCustom(*table, withCache)
		if err != nil {
			return err
		}

		m[table.Name.Source()] = &codeTuple{
			modelCode:       code,
			modelCustomCode: customCode,
		}
	}

	return g.createFile(m)
}
func (g *GormGenerator) StartFromDDL(filename string, withCache, strict bool, database string) error {
	modelList, err := g.genFromDDL(filename, withCache, strict, database)
	if err != nil {
		return err
	}

	return g.createFile(modelList)
}

func (g *GormGenerator) createFile(modelList map[string]*codeTuple) error {
	dirAbs, err := filepath.Abs(g.dir)
	if err != nil {
		return err
	}

	g.dir = dirAbs
	g.pkg = util.SafeString(filepath.Base(dirAbs))
	err = pathx.MkdirIfNotExist(dirAbs)
	if err != nil {
		return err
	}

	for tableName, codes := range modelList {
		tn := stringx.From(tableName)
		modelFilename, err := format.FileNamingFormat(g.cfg.NamingFormat,
			fmt.Sprintf("%s_model", tn.Source()))
		if err != nil {
			return err
		}

		name := util.SafeString(modelFilename) + "_gen.go"
		filename := filepath.Join(dirAbs, name)
		err = os.WriteFile(filename, []byte(codes.modelCode), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// generate error file
	varFilename, err := format.FileNamingFormat(g.cfg.NamingFormat, "vars")
	if err != nil {
		return err
	}

	filename := filepath.Join(dirAbs, varFilename+".go")
	text, err := pathx.LoadTemplate(category, errGormTemplateFile, template.GormError)
	if err != nil {
		return err
	}

	err = util.With("vars").Parse(text).SaveTo(map[string]any{
		"pkg": g.pkg,
	}, filename, false)
	if err != nil {
		return err
	}

	g.Success("Done.")
	return nil
}
func (g *GormGenerator) GenMethod(table Table) (string, error) {
	text, err := pathx.LoadTemplate(category, methodGormTemplateFile, template.GormMethod)
	if err != nil {
		return "", err
	}
	camel := table.Name.ToCamel()
	buffer, err := util.With("create").Parse(text).Execute(map[string]any{
		"data":                  table,
		"upperStartCamelObject": camel,
		"primaryKeyFieldType":   table.PrimaryKey.Field.DataType,
		"primaryKeyField":       table.PrimaryKey.Field.Name.Source(),
		"tableName":             table.FullName.Source(),
	})
	if err != nil {
		return "", err
	}
	buffer.Write([]byte("\n"))

	text1, err := pathx.LoadTemplate(category, findGormOneByFieldTemplateFile, template.GormFindOneByField)
	if err != nil {
		return "", err
	}
	for _, v := range table.UniqueIndex {
		if len(v) == 0 {
			continue
		}
		f := v[0]
		tmpBuf, err := util.With("findOneByField").Parse(text1).Execute(map[string]any{
			"upperStartCamelObject": camel,
			"upperKeyField":         f.Name.ToCamel(),
			"keyFieldType":          f.DataType,
			"keyField":              f.Name.Source(),
		})
		if err != nil {
			continue
		}
		buffer.Write(tmpBuf.Bytes())
		buffer.Write([]byte("\n"))
	}
	return buffer.String(), nil
}
func (g *GormGenerator) genModel(in parser.Table, withCache bool) (string, error) {
	if len(in.PrimaryKey.Name.Source()) == 0 {
		return "", fmt.Errorf("table %s: missing primary key", in.Name.Source())
	}

	primaryKey, uniqueKey := genCacheKeys(in)

	var table Table
	table.Table = in
	table.PrimaryCacheKey = primaryKey
	table.UniqueCacheKey = uniqueKey
	table.ContainsUniqueCacheKey = len(uniqueKey) > 0
	table.ignoreColumns = g.ignoreColumns
	importsCode, err := gormGenImports(in.ContainsTime())
	if err != nil {
		return "", err
	}
	methodCode, err := g.GenMethod(table)
	if err != nil {
		return "", err
	}
	typesCode, err := genGormTypes(table)
	if err != nil {
		return "", err
	}

	code := &gormCode{
		importsCode: importsCode,
		typesCode:   typesCode,
		methodCode:  methodCode,
	}

	output, err := g.executeModel(table, code)
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (g *GormGenerator) genModelCustom(in parser.Table, withCache bool) (string, error) {
	text, err := pathx.LoadTemplate(category, modelCustomTemplateFile, template.ModelCustom)
	if err != nil {
		return "", err
	}

	t := util.With("model-custom").
		Parse(text).
		GoFmt(true)
	output, err := t.Execute(map[string]any{
		"pkg":                   g.pkg,
		"withCache":             withCache,
		"upperStartCamelObject": in.Name.ToCamel(),
		"lowerStartCamelObject": stringx.From(in.Name.ToCamel()).Untitle(),
	})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func (g *GormGenerator) executeModel(table Table, code *gormCode) (*bytes.Buffer, error) {
	text, err := pathx.LoadTemplate(category, modelGormGenTemplateFile, template.GormModelGen)
	if err != nil {

		return nil, err
	}
	t := util.With("model").
		Parse(text).
		GoFmt(true)
	output, err := t.Execute(map[string]interface{}{
		"pkg":     g.pkg,
		"imports": code.importsCode,
		"types":   code.typesCode,
		"method":  code.methodCode,
	})
	if err != nil {
		return nil, err
	}
	return output, nil
}
