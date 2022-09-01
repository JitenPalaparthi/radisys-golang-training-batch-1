package dal

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MAXRETRIES  = 10
	TIMETORETRY = 5
)

func GetConnection(dsn string) (*gorm.DB, error) {
	count := 1
RETRY:
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		if count <= MAXRETRIES {
			time.Sleep(time.Second * time.Duration(TIMETORETRY))
			goto RETRY
		} else {
			return nil, err
		}
	}
	return db, err
}
