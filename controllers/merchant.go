package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hafnisulun/apate-merchants-api/models"
)

type MerchantController struct{}

// GET /merchants
// Find all merchants
func (r MerchantController) FindAll(c *gin.Context) {
	// Bind query
	var query models.FindMerchantsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		log.Println("Bind request query failed:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Get merchants
	var merchants []models.Merchant
	var meta models.Meta = models.Meta{
		Page:  query.Page,
		Limit: query.PerPage,
	}

	tx := models.DB.Debug().Model(&models.Merchant{}).Order("name")

	if query.ResidenceUUID != "" {
		tx.Where("residence_uuid = ?", query.ResidenceUUID)
	}

	tx.Count(&meta.Total)

	tx.Offset((query.Page - 1) * query.PerPage).
		Limit(query.PerPage).
		Find(&merchants)

	meta.Count = len(merchants)

	time.Sleep(3 * time.Second)

	// Send response
	c.JSON(http.StatusOK, gin.H{
		"data": merchants,
		"meta": meta,
	})
}

// GET /merchants/:merchant_uuid
// Find a merchant
func (r MerchantController) FindOne(c *gin.Context) {
	var merchant models.Merchant

	if err := models.DB.
		Where("uuid = ?", c.Param("merchant_uuid")).
		First(&merchant).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merchant})
}

// POST /merchants
// Create a new merchant
func (r MerchantController) Create(c *gin.Context) {
	// Bind input
	var input models.CreateMerchantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create merchant
	merchant := models.Merchant{
		Name:          input.Name,
		Lat:           input.Lat,
		Lon:           input.Lon,
		ResidenceUUID: input.ResidenceUUID,
		ClusterUUID:   input.ClusterUUID,
	}
	models.DB.Create(&merchant)

	// Send response
	c.JSON(http.StatusCreated, gin.H{"data": merchant})
}
