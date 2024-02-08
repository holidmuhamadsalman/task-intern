package controllers

import (
	"net/http"
	"path/filepath"
	"strconv"
	"task-intern-product-api/config"
	"task-intern-product-api/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	file, err := c.FormFile("photo")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed 1", "Message": err.Error()})
		return
	}

	filepath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed 2", "Message": err.Error()})
		return
	}

	hargaStr := c.PostForm("harga")
    harga, err := strconv.Atoi(hargaStr)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed 3", "Message": "Invalid 'harga' value"})
        return
    }

	newProduct := models.Product{
	Nama: c.PostForm("nama"),
	Harga: harga,
	Size: c.PostForm("size"),
	Deskripsi: c.PostForm("deskripsi"),
	Photos: file.Filename,
	}

	if err := config.DB.Create(&newProduct).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Failed 4", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Status": "Success", "Message": "New Product Created", "Data": newProduct})
}

func GetProduct(c *gin.Context) {
	var Products []models.Product
    if err := config.DB.Find(&Products).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Failed", "Message": err.Error()})
        return
    }
 
    c.JSON(http.StatusOK, gin.H{"data": Products})
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	var Product models.Product
	if err := config.DB.First(&Product, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": err.Error()})
			return
	}

	c.JSON(http.StatusOK, gin.H{"data": Product})
}


func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	existingProduct := models.Product{}
	if err := config.DB.First(&existingProduct, id).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": "Product not found"})
			return
	}

	file, err := c.FormFile("photo")
	if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed1", "Message": err.Error()})
			return
	}

	filepath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed2", "Message": err.Error()})
			return
	}

	hargaStr := c.PostForm("harga")
	harga, err := strconv.Atoi(hargaStr)
	if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed3", "Message": "Invalid 'harga' value"})
			return
	}

	updateProduct := models.Product{
			Nama:      c.PostForm("nama"),
			Harga:     harga,
			Size:      c.PostForm("size"),
			Deskripsi: c.PostForm("deskripsi"),
			Photos:  file.Filename,
	}

	if err := config.DB.Model(&existingProduct).Updates(&updateProduct).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Failed4", "Message": err.Error()})
			return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "Message": "Product updated", "Data": updateProduct})
}


func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "Message": "Product deleted"})
}

