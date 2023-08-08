package configs

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	product "apigo/internal/product/models"

	user "apigo/internal/user/models"
)

type Database struct {
	Db *gorm.DB
}

func (d *Database) Connect() error {
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
	if err := d.Db.AutoMigrate(&product.Product{}, &user.User{}); err != nil {
		return err
	}
	return nil
}

func (d *Database) Setup() error {
	if err := d.Connect(); err != nil {
		return err
	}
	if err := d.Migrate(); err != nil {
		return err
	}
	return nil
}
