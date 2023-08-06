package configs

import (
	model "apigo/product/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func (d *Database) Create() {
	//connection string
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=postgres port=%d sslmode=disable",
		EnvConfigs.DB_USER,
		EnvConfigs.DB_PASS,
		EnvConfigs.DB_PORT,
	)
	//connect
	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	d.Db = gormdb

	//migrate model
	d.Db.AutoMigrate(&model.Product{})
}
