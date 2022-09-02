package dal

import (
	"time"

	"github.com/golang/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MAXRETRIES  = 5
	TIMETORETRY = 5
)

func GetConnection(dsn string) (*gorm.DB, error) {
	count := 0
RETRY:
	count++
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if count <= MAXRETRIES {
			glog.Warning("not connected to database:", err.Error())
			time.Sleep(time.Second * time.Duration(TIMETORETRY))
			glog.Info("trying to connect to database:-->", count)
			goto RETRY
		} else {
			return nil, err
		}
	}
	return db, err
}
