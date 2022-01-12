package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/asadbekGo/book-shop-api-gateway/api/handlers/models"
	pb "github.com/asadbekGo/book-shop-api-gateway/genproto/order_service"
	l "github.com/asadbekGo/book-shop-api-gateway/pkg/logger"
	"github.com/asadbekGo/book-shop-api-gateway/pkg/utils"
)

// CreateOrder ...
// @Summary CreateOrder
// @Router /v1/order/ [post]
// @Description This API for creating a new order
// @Tags order
// @Accept  json
// @Produce  json
// @Param Order request body models.Order true "orderCreateRequest"
// @Success 200 {object} models.OrderResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) CreateOrder(c *gin.Context) {
	var (
		body        pb.OrderReq
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

	response, err := h.serviceManager.OrderService().CreateOrder(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create order", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetOrder ...
// @Router /v1/orders/{id} [get]
// @Summary GetOrder
// @Description This API for getting order detail
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.OrderResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) GetOrder(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.OrderService().GetOrder(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get order", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListOrders ...
// @Router /v1/orders [get]
// @Summary ListOrders
// @Description This API for getting list of orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} models.ListOrders
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) ListOrders(c *gin.Context) {
	queryParams := c.Request.URL.Query()

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

	response, err := h.serviceManager.OrderService().ListOrders(
		ctx, &pb.ListReq{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list orders", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateOrder ...
// @Router /v1/order/{id} [put]
// @Summary UpdateOrder
// @Description This API for updating order
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Order request body models.UpdateOrder true "orderUpdateRequest"
// @Success 200 {object} models.OrderResp
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) UpdateOrder(c *gin.Context) {
	var (
		body        pb.OrderReq
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

	response, err := h.serviceManager.OrderService().UpdateOrder(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteOrder ...
// @Router /v1/order/{id} [delete]
// @Summary DeleteOrder
// @Description This API for deleting order
// @Tags order
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} models.Empty
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
func (h *handlerV1) DeleteOrder(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.OrderService().DeleteOrder(
		ctx, &pb.ByIdReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete order", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
