package Config

import (
	"fmt"
	"os"

	"itam/Model/Database"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//func NewDB() *sql.DB {
//	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123456 dbname=DBbarang sslmode=disable")
//	Constant.PanicIfError(err)
//
//	db.SetMaxIdleConns(5)
//	db.SetMaxOpenConns(20)
//	db.SetConnMaxLifetime(60 * time.Minute)
//	db.SetConnMaxIdleTime(10 * time.Minute)
//
//	return db
//}

var (
	DB *gorm.DB
)

func InitDB() {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"))
	con, _ := gorm.Open(postgres.Open(DSN), &gorm.Config{

		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	DB = con
	Migrate(DB)
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&Database.Vendor{},
		&Database.Divisi{},
		&Database.Jabatan{},
		&Database.Role{},
		&Database.Module{},
		&Database.RoleModule{},
		&Database.User{},
		&Database.Asset{},
		&Database.DetaiAsetAplikasi{},
		&Database.DetailAsetHardware{},
		&Database.DetailAsetLisensi{},
		&Database.DetailAsetPerangkat{},
	)
}
