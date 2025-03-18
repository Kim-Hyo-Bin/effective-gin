package databases

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	"effective-gin/configs"
)

var db *gorm.DB

func InitDB(cfg *configs.Config) (*gorm.DB, error) {
	dbConfig := cfg.Database
	var dsn string
	var dialector gorm.Dialector

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
		dbPath := filepath.Join(".", dbConfig.Name+".db") // 프로젝트 루트에 sqlite.db 파일 생성
		dsn = dbPath
		dialector = sqlite.Open(dsn)

		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			file, err := os.Create(dbPath)
			if err != nil {
				return nil, fmt.Errorf("SQLite 데이터베이스 파일 생성 실패: %w", err)
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
		return nil, fmt.Errorf("지원하지 않는 데이터베이스 dialect: %s", dbConfig.Dialect)
	}

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("데이터베이스 연결 실패: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("SQL DB 객체 획득 실패: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	db = gormDB
	return gormDB, nil
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("SQL DB 객체 획득 실패: %w", err)
	}
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("데이터베이스 연결 종료 실패: %w", err)
	}
	return nil
}

func WithContext(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
