package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "apigo/internal/product/models"
)

type Database struct {
	Db *gorm.DB
}

func (d *Database) Create() error {
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
		return err
	}
	d.Db = gormdb

	return nil
}

func (d *Database) Migrate() error {
	//migrate model
	if err := d.Db.AutoMigrate(&model.Product{}); err != nil {
		return err
	}
	return nil
}
