package databases

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"effective-gin/configs"
)

const (
	defaultDBMaxIdleConnections = 10
	defaultDBMaxOpenConntions   = 100
)

var db *gorm.DB

func InitDB(cfg *configs.Config) (*gorm.DB, func(), error) {
	dbConfig := cfg.Database
	var (
		dsn       string
		dialector gorm.Dialector
	)
	switch dbConfig.Dialect {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Name,
		)
		dialector = mysql.Open(dsn)
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
			dbConfig.Host,
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Name,
			dbConfig.Port,
		)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dbPath := filepath.Join(".", dbConfig.Name+".db")
		dsn = dbPath
		dialector = sqlite.Open(dsn)

		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			file, err := os.Create(dbPath)
			if err != nil {
				return nil, nil, fmt.Errorf("failed generate sqlite file: %w", err)
			}
			file.Close()
		}

	case "sqlserver":
		dsn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			dbConfig.User,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Name,
		)
		dialector = sqlserver.Open(dsn)

	default:
		return nil, nil, fmt.Errorf("unsupported database dialect: %s", dbConfig.Dialect)
	}

	//TODO: seperate dialector for each database

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed connect database: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("failed get objacet DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(defaultDBMaxIdleConnections)
	sqlDB.SetMaxOpenConns(defaultDBMaxOpenConntions)

	db = gormDB
	cleanup := func() {
		if err := sqlDB.Close(); err != nil {
			return
		}
	}
	return gormDB, cleanup, nil
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("failed to get database connection")
	}
	return db
}

func WithContext(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
