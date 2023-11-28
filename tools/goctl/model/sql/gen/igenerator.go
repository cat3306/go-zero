package gen

import "github.com/zeromicro/go-zero/tools/goctl/model/sql/model"

type Generator interface {
	StartFromDDL(filename string, withCache, strict bool, database string) error
	StartFromInformationSchema(tables map[string]*model.Table, withCache, strict bool) error
}
