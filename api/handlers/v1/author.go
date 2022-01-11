package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/asadbekGo/book-shop-api-gateway/api/docs" // swag nolint:errcheck
	pb "github.com/asadbekGo/book-shop-api-gateway/genproto/catalog_service"
	l "github.com/asadbekGo/book-shop-api-gateway/pkg/logger"
	"github.com/asadbekGo/book-shop-api-gateway/pkg/utils"
)

// CreateAuthor ...
// @Summary CreateAuthor
// @Router /v1/author/ [post]
// @Description This API for creating a new author
// @Tags author
// @Accept  json
// @Produce  json
// @Param Author request body models.Author true "authorCreateRequest"
// @Success 200 {object} models.AuthorResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) CreateAuthor(c *gin.Context) {
	var (
		body        pb.Author
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

	response, err := h.serviceManager.CatalogService().CreateAuthor(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create author", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetAuthor ...
// @Router /v1/authors/{id} [get]
// @Summary GetAuthor
// @Description This API for getting author ListAuthors
// @Tags author
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.AuthorResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) GetAuthor(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().GetAuthor(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get author", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAuthors ...
// @Router /v1/authors [get]
// @Summary GetAuthors
// @Description This API for getting list of authors
// @Tags author
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListAuthors
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) GetAuthors(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	fmt.Println("OK")

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().GetAuthors(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list authors", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateAuthor ...
// @Router /v1/author/{id} [put]
// @Summary UpdateAuthor
// @Description This API for updating author
// @Tags author
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Author request body models.UpdateAuthor true "authorUpdateRequest"
// @Success 200 {object} models.AuthorResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) UpdateAuthor(c *gin.Context) {
	var (
		body        pb.Author
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
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().UpdateAuthor(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update author", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteAuthor ...
// @Router /v1/author/{id} [delete]
// @Summary DeleteAuthor
// @Description This API for deleting author
// @Tags author
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DeleteAuthor(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CatalogService().DeleteAuthor(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete author", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
