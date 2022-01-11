package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/asadbekGo/book-shop-api-gateway/api/handlers/models"
	pb "github.com/asadbekGo/book-shop-api-gateway/genproto/catalog_service"
	l "github.com/asadbekGo/book-shop-api-gateway/pkg/logger"
)

// CreateBookCategory ...
// @Summary CreateBookCategory
// @Router /v1/bookCategory/ [post]
// @Description This API for creating a new bookCategory
// @Tags bookCategory
// @Accept  json
// @Produce  json
// @Param BookCategory request body models.BookCategory true "bookCategoryCreateRequest"
// @Success 200 {object} models.BookResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) CreateBookCategory(c *gin.Context) {
	var (
		body        pb.BookCategory
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().CreateBookCategory(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create bookCategory", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// DeleteBookCategory ...
// @Router /v1/bookCategory/{id} [delete]
// @Summary DeleteBookCategory
// @Description This API for deleting bookCategory
// @Tags bookCategory
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param BookCategory request body models.BookCategoryDelete true "bookCategoryCreateRequest"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DeleteBookCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	var body struct {
		CategoryId string `json:"categoryId"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	fmt.Println(body)

	guid1 := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().DeleteBookCategory(
		ctx, &pb.BookCategory{
			BookId:     guid1,
			CategoryId: body.CategoryId,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete bookCategory", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
