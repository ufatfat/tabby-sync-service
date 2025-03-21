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
		log.Printf("dsn: %s", dsn)
		if db, err = gorm.Open(
			mysql.Open(dsn),
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

	if err := db.AutoMigrate(&tables.Config{}, &tables.User{}); err != nil {
		panic("Database error: " + err.Error())
	}
}
