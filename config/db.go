package configs

import (
	"apigo/product"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func (d *Database) Create() {
	fmt.Println(EnvConfigs.DB_NAME)
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=postgres port=%d sslmode=disable", EnvConfigs.DB_USER, EnvConfigs.DB_PASS, EnvConfigs.DB_PORT)
	fmt.Println(dsn)

	gormdb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	d.Db = gormdb
	d.Db.AutoMigrate(&product.Product{})
}
