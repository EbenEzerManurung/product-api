package seeder

import (
    "log"

    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"

    "product-api/model"
)

func RunAll(db *gorm.DB) {
    log.Println("🌱 Menjalankan seeder...")

    seedUsers(db)
    seedProducts(db)

    log.Println("✅ Seeder selesai")
}

func seedUsers(db *gorm.DB) {
    var count int64
    db.Model(&model.User{}).Count(&count)
    if count > 0 {
        log.Println("ℹ️  User sudah ada, seeder dilewati")
        return
    }

    hashed, _ := bcrypt.GenerateFromPassword([]byte("supersecret"), bcrypt.DefaultCost)
    users := []model.User{
        {
            Username: "jhon_doe",
            Password: string(hashed),
        },
    }

    if err := db.Create(&users).Error; err != nil {
        log.Printf("⚠️  Gagal seed user: %v", err)
    } else {
        log.Printf("✅ Berhasil seed %d user", len(users))
    }
}

func seedProducts(db *gorm.DB) {
    var count int64
    db.Model(&model.Product{}).Count(&count)
    if count > 0 {
        log.Println("ℹ️  Produk sudah ada, seeder dilewati")
        return
    }

    products := []model.Product{
        {
            Title:       "Awesome T-Shirt",
            Price:       99.99,
            Description: "High-quality cotton t-shirt",
            Category:    "Clothes",
            Images:      model.JSONStringSlice{"https://placeimg.com/640/480/any"},
            CreatedBy:   "Jhon Doe",
            CreatedByID: "1",
            UpdatedBy:   "Jhon Doe",
            UpdatedByID: "1",
        },
        {
            Title:       "Wireless Headphones",
            Price:       199.50,
            Description: "Noise-cancelling over-ear headphones",
            Category:    "Electronics",
            Images:      model.JSONStringSlice{"https://placeimg.com/640/480/tech"},
            CreatedBy:   "Jhon Doe",
            CreatedByID: "1",
            UpdatedBy:   "Jhon Doe",
            UpdatedByID: "1",
        },
        {
            Title:       "Leather Wallet",
            Price:       45.00,
            Description: "Genuine leather bifold wallet",
            Category:    "Accessories",
            Images:      model.JSONStringSlice{"https://placeimg.com/640/480/fashion"},
            CreatedBy:   "Jhon Doe",
            CreatedByID: "1",
            UpdatedBy:   "Jhon Doe",
            UpdatedByID: "1",
        },
    }

    if err := db.Create(&products).Error; err != nil {
        log.Printf("⚠️  Gagal seed produk: %v", err)
    } else {
        log.Printf("✅ Berhasil seed %d produk", len(products))
    }
}