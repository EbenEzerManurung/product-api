package database

import (
    "fmt"
    "log"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"

    "product-api/config"
    "product-api/model"
)

func Connect(cfg *config.Config) *gorm.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatalf("Gagal terhubung ke database: %v", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Gagal mengambil sql.DB: %v", err)
    }

    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

    log.Println("✅ Database terhubung")
    return db
}

// FreshMigrate menghapus semua tabel lalu membuat ulang berdasarkan model.
func FreshMigrate(db *gorm.DB) {
    db.Exec("SET FOREIGN_KEY_CHECKS = 0")

    tables := []string{"products", "users"}
    for _, table := range tables {
        if err := db.Migrator().DropTable(table); err != nil {
            log.Printf("⚠️  Peringatan: gagal drop tabel %s: %v", table, err)
        }
    }

    db.Exec("SET FOREIGN_KEY_CHECKS = 1")

    if err := db.AutoMigrate(&model.User{}, &model.Product{}); err != nil {
        log.Fatalf("❌ Migrasi gagal: %v", err)
    }

    log.Println("✅ Tabel berhasil dimigrasikan")
}