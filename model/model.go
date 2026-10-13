package model

import (
    "database/sql/driver"
    "encoding/json"
    "errors"
    "time"

    "gorm.io/gorm"
)

type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Username  string         `gorm:"uniqueIndex;size:255;not null" json:"username"`
    Password  string         `gorm:"size:255;not null" json:"-"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Product struct {
    ID          uint            `gorm:"primaryKey" json:"id"`
    Title       string          `gorm:"size:255;not null" json:"title"`
    Price       float64         `gorm:"not null" json:"price"`
    Description string          `gorm:"type:text" json:"description"`
    Category    string          `gorm:"size:100;not null;index" json:"category"`
    Images      JSONStringSlice `gorm:"type:json" json:"images"`
    CreatedAt   time.Time       `json:"created_at"`
    CreatedBy   string          `gorm:"size:255" json:"created_by"`
    CreatedByID string          `gorm:"size:50" json:"created_by_id"`
    UpdatedAt   time.Time       `json:"updated_at"`
    UpdatedBy   string          `gorm:"size:255" json:"updated_by"`
    UpdatedByID string          `gorm:"size:50" json:"updated_by_id"`
    DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"`
}

// JSONStringSlice menyimpan []string sebagai JSON di MySQL.
type JSONStringSlice []string

func (j JSONStringSlice) Value() (driver.Value, error) {
    if j == nil {
        return nil, nil
    }
    return json.Marshal(j)
}

func (j *JSONStringSlice) Scan(value interface{}) error {
    if value == nil {
        *j = nil
        return nil
    }
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("gagal scan JSONStringSlice")
    }
    return json.Unmarshal(bytes, j)
}