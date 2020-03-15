package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegDB(t *testing.T) {
	pgConfig := "host=127.0.0.1  port=5432 user=pguser dbname=testdb sslmode=disable password=pgpass"
	RegDB("default", DB_POSTGRES, pgConfig)
	msConfig := "host=127.0.0.1  port=5432 user=pguser dbname=testdb sslmode=disable password=pgpass"
	RegDB("Mysql", DB_MYSQL, msConfig)

	config, e := dbConfigs["default"]
	assert.Equal(t, e, true)
	assert.Equal(t, config.DbType, DB_POSTGRES)
	assert.Equal(t, config.Connect, pgConfig)
	config2, e2 := dbConfigs["Mysql"]
	assert.Equal(t, e2, true)
	assert.Equal(t, config2.DbType, DB_MYSQL)
	assert.Equal(t, config2.Connect, msConfig)
}
func TestGetDB(t *testing.T) {
	RegDB("default", DB_POSTGRES, "host=127.0.0.1  port=5432 user=pguser dbname=testdb sslmode=disable password=pgpass")
	db := GetDB("default")
	isExist := db.HasTable("userasinfo")
	assert.Equal(t, isExist, false)
}
