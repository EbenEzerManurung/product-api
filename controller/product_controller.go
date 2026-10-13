package controller

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"

    "product-api/service"
)

type ProductController struct {
    productService *service.ProductService
}

func NewProductController(s *service.ProductService) *ProductController {
    return &ProductController{productService: s}
}


func (h *ProductController) GetAll(c *gin.Context) {
    search := c.Query("search")
    category := c.Query("category")
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

    products, total, err := h.productService.GetAll(search, category, limit, page)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  products,
        "total": total,
        "page":  page,
        "limit": limit,
    })
}


func (h *ProductController) GetByID(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk tidak valid"})
        return
    }

    product, err := h.productService.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Produk tidak ditemukan"})
        return
    }

    c.JSON(http.StatusOK, product)
}


func (h *ProductController) Create(c *gin.Context) {
    var input service.ProductInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    username, _ := c.Get("username")
    usernameStr, _ := username.(string)
    userID := "1"

    product, err := h.productService.Create(input, usernameStr, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, product)
}


func (h *ProductController) Update(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk tidak valid"})
        return
    }

    var input service.ProductInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    username, _ := c.Get("username")
    usernameStr, _ := username.(string)
    userID := "1"

    product, err := h.productService.Update(uint(id), input, usernameStr, userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, product)
}


func (h *ProductController) Delete(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID produk tidak valid"})
        return
    }

    if err := h.productService.Delete(uint(id)); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Produk berhasil dihapus"})
}