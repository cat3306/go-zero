package gen

import (
	"fmt"

	"github.com/zeromicro/go-zero/tools/goctl/model/sql/template"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	category                              = "model"
	deleteTemplateFile                    = "delete.tpl"
	deleteMethodTemplateFile              = "interface-delete.tpl"
	fieldTemplateFile                     = "field.tpl"
	findOneTemplateFile                   = "find-one.tpl"
	findOneMethodTemplateFile             = "interface-find-one.tpl"
	findOneByFieldTemplateFile            = "find-one-by-field.tpl"
	findOneByFieldMethodTemplateFile      = "interface-find-one-by-field.tpl"
	findOneByFieldExtraMethodTemplateFile = "find-one-by-field-extra-method.tpl"
	importsTemplateFile                   = "import.tpl"
	importsWithNoCacheTemplateFile        = "import-no-cache.tpl"
	insertTemplateFile                    = "insert.tpl"
	insertTemplateMethodFile              = "interface-insert.tpl"
	modelGenTemplateFile                  = "model-gen.tpl"
	modelCustomTemplateFile               = "model.tpl"
	modelNewTemplateFile                  = "model-new.tpl"
	tableNameTemplateFile                 = "table-name.tpl"
	tagTemplateFile                       = "tag.tpl"
	typesTemplateFile                     = "types.tpl"
	updateTemplateFile                    = "update.tpl"
	updateMethodTemplateFile              = "interface-update.tpl"
	varTemplateFile                       = "var.tpl"
	errTemplateFile                       = "err.tpl"
	tagGormTemplateFile                   = "gorm-tag.tpl"
	importGormTemplateFile                = "gorm-import.tpl"
	typesGormTemplateFile                 = "gorm-types.tpl"
	modelGormGenTemplateFile              = "gorm-model-gen.tpl"
	errGormTemplateFile                   = "gorm-err.tpl"
	methodGormTemplateFile                = "gorm-method.tpl"
	findGormOneByFieldTemplateFile        = "gorm-find-one-by-field.tpl"
	modelGormCustomTemplateFile           = "gorm-custom-model.tpl"
)

var templates = map[string]string{
	deleteTemplateFile:                    template.Delete,
	deleteMethodTemplateFile:              template.DeleteMethod,
	fieldTemplateFile:                     template.Field,
	findOneTemplateFile:                   template.FindOne,
	findOneMethodTemplateFile:             template.FindOneMethod,
	findOneByFieldTemplateFile:            template.FindOneByField,
	findOneByFieldMethodTemplateFile:      template.FindOneByFieldMethod,
	findOneByFieldExtraMethodTemplateFile: template.FindOneByFieldExtraMethod,
	importsTemplateFile:                   template.Imports,
	importsWithNoCacheTemplateFile:        template.ImportsNoCache,
	insertTemplateFile:                    template.Insert,
	insertTemplateMethodFile:              template.InsertMethod,
	modelGenTemplateFile:                  template.ModelGen,
	modelCustomTemplateFile:               template.ModelCustom,
	modelNewTemplateFile:                  template.New,
	tableNameTemplateFile:                 template.TableName,
	tagTemplateFile:                       template.Tag,
	typesTemplateFile:                     template.Types,
	updateTemplateFile:                    template.Update,
	updateMethodTemplateFile:              template.UpdateMethod,
	varTemplateFile:                       template.Vars,
	errTemplateFile:                       template.Error,
	tagGormTemplateFile:                   template.GormTag,
	importGormTemplateFile:                template.GormImports,
	modelGormGenTemplateFile:              template.GormModelGen,
	errGormTemplateFile:                   template.GormError,
	methodGormTemplateFile:                template.GormMethod,
	findGormOneByFieldTemplateFile:        template.GormFindOneByField,
	typesGormTemplateFile:                 template.GormTypes,
	modelGormCustomTemplateFile:           template.GormModelCustom,
}

// Category returns model const value
func Category() string {
	return category
}

// Clean deletes all template files
func Clean() error {
	return pathx.Clean(category)
}

// GenTemplates creates template files if not exists
func GenTemplates() error {
	return pathx.InitTemplates(category, templates)
}

// RevertTemplate reverts the deleted template files
func RevertTemplate(name string) error {
	content, ok := templates[name]
	if !ok {
		return fmt.Errorf("%s: no such file name", name)
	}

	return pathx.CreateTemplate(category, name, content)
}

// Update provides template clean and init
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return pathx.InitTemplates(category, templates)
}
