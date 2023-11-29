package {{.pkgName}}
import (
	"fmt"
	"{{.configPkg}}"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
    {{.exportVar}} *gorm.DB
)

func {{.initCode}}() error {
    mysqlConn := fmt.Sprintf("%s:%s@(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConf.DbConfig.User,
		config.AppConf.DbConfig.Pwd,
		config.AppConf.DbConfig.Ip,
		config.AppConf.DbConfig.Port)
	myDb, err := gorm.Open(mysql.Open(mysqlConn))
	if err != nil {
		return err
	}
	sqlDb, err := myDb.DB()
	if err != nil {
		return err
	}
	err = sqlDb.Ping()
	if err != nil {
		return err
	}
	if true {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color

			},
		)
		myDb.Logger = newLogger
	}
	DB = myDb
	return err
}