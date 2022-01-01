package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Host       string
	User       string
	Password   string
	DbName     string
	Port       string
	SslMode    string
	TimeZone   string
	Connection *gorm.DB
}

func NewProductionDB() *DB {
	c := NewConfig()
	return NewProductionDB(&DB{
		Host:       c.DB.Production.Host,
		User:       c.DB.Production.User,
		Password:   c.DB.Production.Password,
		DbName:     c.DB.Production.DbName,
		Port:       c.DB.Production.Port,
		SslMode:    c.DB.Production.SslMode,
		TimeZone:   c.DB.Production.TimeZone,
		Connection: &gorm.DB{},
	})
}

func NewDevelopmentDB() *DB {
	c := NewConfig()
	return NewDevelopmentDB(&DB{
		Host:       c.DB.Development.Host,
		User:       c.DB.Development.User,
		Password:   c.DB.Development.Password,
		DbName:     c.DB.Development.DbName,
		Port:       c.DB.Development.Port,
		SslMode:    c.DB.Development.SslMode,
		TimeZone:   c.DB.Development.TimeZone,
		Connection: &gorm.DB{},
	})
}

func NewDB(d *DB) *DB {
	dsn := "host=" + d.Host + " user=" + d.User + " password=" + d.Password + " dbname=" + d.DbName + " port=" + d.Port + " sslmode=" + d.SslMode + " TimeZone=" + d.TimeZone
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	d.Connection = db
	return d
}

func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
