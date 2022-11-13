package helpers

import "github.com/gin-gonic/gin"

func GetQueryData(c *gin.Context, query string) string {
	return c.Query(query)
}
