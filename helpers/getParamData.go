package helpers

import "github.com/gin-gonic/gin"

func GetParamData(c *gin.Context, param string) string {
	return c.Param(param)
}
