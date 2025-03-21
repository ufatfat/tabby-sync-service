package services

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"tabby-sync/configs"
	"tabby-sync/consts"
	"tabby-sync/tables"
	"time"
)

var db *gorm.DB

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), consts.DatabaseConnectTimeout)

	go func(ctx context.Context) {
		var err error
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai",
			configs.Database.Username,
			configs.Database.Password,
			configs.Database.Host,
			configs.Database.Port,
			configs.Database.Database,
		)
		if db, err = gorm.Open(
			mysql.New(mysql.Config{
				DSN:                  dsn,
				DisableWithReturning: true,
			}),
			&gorm.Config{
				Logger: newLogger,
			},
		); err != nil {
			panic("Database connect error: " + err.Error())
		}

		sqlDB, err := db.DB()
		if err != nil {
			panic("Database error")
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(10000)
		sqlDB.SetConnMaxLifetime(time.Second * 3)

		cancel()
	}(ctx)

	select {
	case <-ctx.Done():
		switch err := ctx.Err(); {
		case errors.Is(err, context.DeadlineExceeded):
			panic("Database connection initialization timed out.")
		case errors.Is(err, context.Canceled):
			fmt.Println("Database successfully connected.")
		}
	}

	migrate()
}

func migrate() {
	_, err := os.Open("db_migrate.lock")
	if err == nil {
		return
	}
	if !os.IsNotExist(err) {
		panic(err)
	}
	_, err = os.Create("db_migrate.lock")
	if err != nil {
		panic(err)
	}

	if err = db.AutoMigrate(&tables.Config{}, &tables.User{}, &tables.OAuthBinding{}); err != nil {
		panic("Database error: " + err.Error())
	}

	createAdmin()
}

func createAdmin() {
	if id, err := NewUser(0, "", configs.Admin.Username, configs.Admin.Email, "", ""); err != nil {
		panic(err)
	} else {
		log.Printf("Created admin user: %d\n", id)
	}
}
