package config

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/xerrors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetEnv(envDirPath string) error {
	// 実行環境取得
	env := os.Getenv("ENV")
	if env != "dev" && env != "stg" && env != "prod" {
		return xerrors.Errorf("環境変数ENVの値が正しく設定されていません. env = %s", env)
	}

	// 環境変数取得
	if err := godotenv.Load(envDirPath + "/" + env + ".env"); err != nil {
		return xerrors.Errorf("Failed to load %s. : %w", envDirPath, err)
	}

	return nil
}

func createDsn() string {
	return os.Getenv("DB_USER") +
		":" + os.Getenv("DB_PASS") +
		"@" + os.Getenv("DB_PROTOCOL") +
		"(" + os.Getenv("DB_ADDRESS") +
		":" + os.Getenv("DB_PORT") + ")" +
		"/" + os.Getenv("DB_NAME")
}

// DB接続
func ConnectDB() (*gorm.DB, error) {
	// DB接続
	dsn := createDsn()

	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, xerrors.Errorf("Failed to initialize db session. dsn = %s : %w", dsn, err)
	}

	return dbConn, nil
}
