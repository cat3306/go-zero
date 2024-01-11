package {{.pkg}}

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func IgnoreRecordNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return err
}

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
