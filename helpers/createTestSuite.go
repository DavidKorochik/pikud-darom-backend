package helpers

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CreateTestSuite(t *testing.T) (*assert.Assertions, *httptest.ResponseRecorder, *gin.Engine) {
	a := assert.New(t)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	return a, w, r
}
