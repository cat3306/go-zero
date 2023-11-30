package gogen

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/format"
	"github.com/zeromicro/go-zero/tools/goctl/vars"
)

const (
	configFile     = "config"
	configTypeFile = "types"
	jwtTemplate    = ` struct {
		AccessSecret string
		AccessExpire int64
	}
`
	jwtTransTemplate = ` struct {
		Secret     string
		PrevSecret string
	}
`
)

//go:embed config.tpl
var configTemplate string

//go:embed config-type.tpl
var configTypeTemplate string

func genConfig(dir string, cfg *config.Config, api *spec.ApiSpec, component string) error {
	filename, err := format.FileNamingFormat(cfg.NamingFormat, configFile)
	if err != nil {
		return err
	}

	authNames := getAuths(api)
	var auths []string
	for _, item := range authNames {
		auths = append(auths, fmt.Sprintf("%s %s", item, jwtTemplate))
	}

	jwtTransNames := getJwtTrans(api)
	var jwtTransList []string
	for _, item := range jwtTransNames {
		jwtTransList = append(jwtTransList, fmt.Sprintf("%s %s", item, jwtTransTemplate))
	}
	authImportStr := fmt.Sprintf("\"%s/rest\"", vars.ProjectOpenSourceURL)
	data := genFileDataMap(component)
	data["authImport"] = authImportStr
	data["auth"] = strings.Join(auths, "\n")
	data["jwtTrans"] = strings.Join(jwtTransList, "\n")
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          configDir,
		filename:        filename + ".go",
		templateName:    "configTemplate",
		category:        category,
		templateFile:    configTemplateFile,
		builtinTemplate: configTemplate,
		data:            data,
	})
}
func genFileDataMap(component string) map[string]any {
	r := make(map[string]any)
	if component == "" {
		return r
	}
	list := strings.Split(component, ",")
	for _, v := range list {
		_, ok := componentsMap[v]
		r[v] = ok
	}
	return r
}
func genComponentConfigType(dir string, cfg *config.Config, api *spec.ApiSpec, component string) error {
	if component == "" {
		return nil
	}
	filename, err := format.FileNamingFormat(cfg.NamingFormat, configTypeFile)
	if err != nil {
		return err
	}
	data := genFileDataMap(component)
	data["pkgName"] = configPackage
	return genFile(fileGenConfig{
		dir:             dir,
		subdir:          configDir,
		filename:        filename + ".go",
		templateName:    "configTypeTemplate",
		category:        category,
		templateFile:    configTypeTemplateFile,
		builtinTemplate: configTypeTemplate,
		data:            data,
	})
}
