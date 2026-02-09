package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
)

func (h *Handler) Login(c *gin.Context) {
	var request memberships.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	accessToken, err := h.service.Login(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, memberships.LoginResponse{
		AccessToken: accessToken,
	})

}
