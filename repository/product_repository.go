package repository

import (
    "gorm.io/gorm"

    "product-api/model"
)

type ProductRepository struct {
    DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
    return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *model.Product) error {
    return r.DB.Create(product).Error
}

func (r *ProductRepository) FindByID(id uint) (*model.Product, error) {
    var product model.Product
    err := r.DB.First(&product, id).Error
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepository) FindAll(search, category string, limit, offset int) ([]model.Product, int64, error) {
    var products []model.Product
    var total int64

    query := r.DB.Model(&model.Product{})

    if search != "" {
        query = query.Where("title LIKE ?", "%"+search+"%")
    }
    if category != "" {
        query = query.Where("category = ?", category)
    }

    query.Count(&total)

    err := query.Limit(limit).Offset(offset).Order("id DESC").Find(&products).Error
    return products, total, err
}

func (r *ProductRepository) Update(product *model.Product) error {
    return r.DB.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
    return r.DB.Delete(&model.Product{}, id).Error
}