package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisplayErrorMsg(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
