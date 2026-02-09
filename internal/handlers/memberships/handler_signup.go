package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
)

func (h *Handler) SignUp(c *gin.Context) {
	var request memberships.SignUpRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err := h.service.SignUp(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
