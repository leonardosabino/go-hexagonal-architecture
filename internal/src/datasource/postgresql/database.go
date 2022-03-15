package datasource

import (
	"context"
	"log"
	"sync"

	"hexagonal/template/internal/src/config"

	"github.com/joomcode/errorx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	runOnceConnect sync.Once
	database       *Database
)

func GetDatabase() Database {
	runOnceConnect.Do(func() {
		database = &Database{
			connection: nil,
			error:      nil,
		}
	})

	if database.connection == nil {
		database.Connect()
	}

	return *database
}

type Database struct {
	connection *gorm.DB
	error      error
}

func (c *Database) Connect() {
	dsn := config.GetConfig().DatabaseDSN
	connection, error := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if error != nil {
		log.Println(errorx.Decorate(error, "Failed to connect in database"))
	}
	c.connection = connection
	c.error = error
}

func (c Database) Ping(ctx context.Context) *errorx.Error {
	results, error := c.connection.DB()
	if error != nil {
		return errorx.Decorate(error, "Unable to get connection")
	}

	if pingError := results.Ping(); pingError != nil {
		return errorx.Decorate(error, "Unable to ping")
	}
	return nil
}
