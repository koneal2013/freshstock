package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/koneal2013/freshstock/internal/model"
	"github.com/koneal2013/freshstock/internal/store"
)

type Handlers struct {
	store *store.ProduceStore
}

func NewHandlers(store *store.ProduceStore) *Handlers {
	return &Handlers{store: store}
}

func (h *Handlers) AddProduce(c *gin.Context) {
	var produce model.Produce
	if err := c.ShouldBindJSON(&produce); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.store.AddProduce(&produce); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, produce)
}

func (h *Handlers) GetProduceByCode(c *gin.Context) {
	code := c.Param("code")

	produce, err := h.store.GetProduceByCode(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, produce)
}

func (h *Handlers) SearchProduce(c *gin.Context) {
	query := c.Query("q")

	results, err := h.store.SearchProduce(query)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func (h *Handlers) DeleteProduce(c *gin.Context) {
	code := c.Param("code")

	if err := h.store.DeleteProduce(code); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
