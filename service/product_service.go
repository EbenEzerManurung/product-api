package service

import (
    "errors"
    "time"

    "product-api/model"
    "product-api/repository"
)

type ProductService struct {
    repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

type ProductInput struct {
    Title       string   `json:"title" binding:"required"`
    Price       float64  `json:"price" binding:"required,gt=0"`
    Description string   `json:"description"`
    Category    string   `json:"category" binding:"required"`
    Images      []string `json:"images" binding:"required,min=1"`
}

func (s *ProductService) Create(input ProductInput, username string, userID string) (*model.Product, error) {
    product := &model.Product{
        Title:       input.Title,
        Price:       input.Price,
        Description: input.Description,
        Category:    input.Category,
        Images:      model.JSONStringSlice(input.Images),
        CreatedBy:   username,
        CreatedByID: userID,
        UpdatedBy:   username,
        UpdatedByID: userID,
    }
    if err := s.repo.Create(product); err != nil {
        return nil, err
    }
    return product, nil
}

func (s *ProductService) GetByID(id uint) (*model.Product, error) {
    return s.repo.FindByID(id)
}

func (s *ProductService) GetAll(search, category string, limit, page int) ([]model.Product, int64, error) {
    if limit <= 0 {
        limit = 10
    }
    if page <= 0 {
        page = 1
    }
    offset := (page - 1) * limit
    return s.repo.FindAll(search, category, limit, offset)
}

func (s *ProductService) Update(id uint, input ProductInput, username string, userID string) (*model.Product, error) {
    product, err := s.repo.FindByID(id)
    if err != nil {
        return nil, errors.New("produk tidak ditemukan")
    }

    product.Title = input.Title
    product.Price = input.Price
    product.Description = input.Description
    product.Category = input.Category
    product.Images = model.JSONStringSlice(input.Images)
    product.UpdatedBy = username
    product.UpdatedByID = userID
    product.UpdatedAt = time.Now()

    if err := s.repo.Update(product); err != nil {
        return nil, err
    }
    return product, nil
}

func (s *ProductService) Delete(id uint) error {
    _, err := s.repo.FindByID(id)
    if err != nil {
        return errors.New("produk tidak ditemukan")
    }
    return s.repo.Delete(id)
}