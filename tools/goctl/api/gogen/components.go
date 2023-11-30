package gogen

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"strings"
)

const (
	initComponent        = "initComponent"
	mysqlComponent       = "mysql"
	redisComponent       = "redis"
	componentPackageName = "component"
	componentFile        = "component.go"
)

type componentStruct struct {
	fileName        string
	templateName    string
	templateFile    string
	data            map[string]any
	builtinTemplate string
	initFuncCode    string
	exportVar       string
}

var (
	//go:embed gorm-mysql.tpl
	gormMysqlTemplate string

	//go:embed redis.tpl
	redisTemplate string
	//go:embed component.tpl
	componentTemplate string
	componentsMap     = map[string]componentStruct{
		initComponent: {
			fileName:     "component.go",
			templateName: "componentTemplate",
			templateFile: "component.tpl",
			data: map[string]any{
				"pkgName": componentPackageName,
			},
			builtinTemplate: componentTemplate,
		},
		mysqlComponent: {
			initFuncCode: "initMysql",
			fileName:     "mysql.go",
			templateName: "mysqlTemplate",
			templateFile: "gorm-mysql.tpl",
			exportVar:    "DB",
			data: map[string]any{
				"pkgName":   componentPackageName,
				"initCode":  "initMysql",
				"exportVar": "DB",
			},
			builtinTemplate: gormMysqlTemplate,
		},
		redisComponent: {
			fileName:        "redis.go",
			templateName:    "redisTemplate",
			templateFile:    "redis.tpl",
			data:            map[string]any{},
			builtinTemplate: redisTemplate,
		},
	}
)

func CheckComponentsValid(components string) ([]string, error) {
	if components == "" {
		return nil, nil
	}
	list := strings.Split(components, ",")
	for _, v := range list {
		_, ok := componentsMap[v]
		if !ok {
			return nil, fmt.Errorf("unknown component %s", v)
		}
	}
	return list, nil
}
func genComponents(dir string, rootPkg string, cfg *config.Config, api *spec.ApiSpec, components string) error {
	if components == "" {
		return nil
	}
	cs, err := CheckComponentsValid(components)
	if err != nil {
		return err
	}
	initFunc := func(code string, init string) string {
		if init == "" {
			return code
		}
		code += fmt.Sprintf(`
		err:=%s ()
		if err !=nil{
			return err
		}
`, init)
		return code
	}
	initCode := ""
	//cs = append(cs, initComponent)
	for _, c := range cs {
		conf, ok := componentsMap[c]
		if !ok {
			continue
		}
		conf.data["configPkg"] = pathx.JoinPackages(rootPkg, configDir)
		err = genFile(fileGenConfig{
			dir:             dir,
			subdir:          componentDir,
			filename:        conf.fileName,
			templateName:    conf.templateName,
			category:        category,
			templateFile:    conf.templateFile,
			builtinTemplate: conf.builtinTemplate,
			data:            conf.data,
		})
		if err != nil {
			return err
		}
		initCode = initFunc(initCode, conf.initFuncCode)
	}
	tmpConf := componentsMap[initComponent]
	tmpConf.data["initCode"] = initCode
	err = genFile(fileGenConfig{
		dir:             dir,
		subdir:          componentDir,
		filename:        tmpConf.fileName,
		templateName:    tmpConf.templateName,
		category:        category,
		templateFile:    tmpConf.templateFile,
		builtinTemplate: tmpConf.builtinTemplate,
		data:            tmpConf.data,
	})
	return err
}
