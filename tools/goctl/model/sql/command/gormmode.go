package command

import (
	"errors"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/gen"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"strings"
)

func fromGormDDL(arg ddlArg) error {
	log := console.NewConsole(arg.idea)
	src := strings.TrimSpace(arg.src)
	if len(src) == 0 {
		return errors.New("expected path or path globbing patterns, but nothing found")
	}

	files, err := util.MatchFiles(src)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errNotMatched
	}

	generator, err := gen.NewGormGenerator(arg.dir, arg.cfg, gen.WithConsoleOption(log))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = generator.StartFromDDL(file, arg.cache, arg.strict, arg.database)
		if err != nil {
			return err
		}
	}
	return nil
}
