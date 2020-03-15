package database

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB_TYPE string

const (
	DB_MYSQL    DB_TYPE = "mysql"
	DB_POSTGRES DB_TYPE = "postgres"
)

type DataBaseConfig struct {
	DbType  DB_TYPE
	Connect string
}

var dbLock sync.Mutex
var dbIns map[DB_TYPE]*gorm.DB
var dbConfigs map[string]*DataBaseConfig
var dbOnce sync.Once
var dbConfigOnce sync.Once

func RegDB(key string, dbType DB_TYPE, connect string) {
	dbConfigOnce.Do(func() {
		dbConfigs = make(map[string]*DataBaseConfig)
	})
	dbConfigs[key] = &DataBaseConfig{DbType: dbType, Connect: connect}
}

func GetDB(key string) *gorm.DB {
	config, exist := dbConfigs[key]
	if !exist {
		panic("not found DBConfig")
	}
	dbOnce.Do(func() {
		dbIns = make(map[DB_TYPE]*gorm.DB)
	})
	dbLock.Lock()
	defer dbLock.Unlock()
	if _, isExist := dbIns[config.DbType]; !isExist {
		db, err := gorm.Open(string(config.DbType), config.Connect)
		if err != nil {
			panic("connect database fail")
		}
		defer db.Close()
		db.SingularTable(true)
		dbIns[config.DbType] = db
	}
	return dbIns[config.DbType]
}
