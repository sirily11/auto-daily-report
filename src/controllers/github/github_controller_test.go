package github

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestWebhook(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//t.Run("it should return 200 when the JSON is valid", func(t *testing.T) {
	//	router := gin.Default()
	//	controller := NewController()
	//	controller.RegisterRoutes(router.Group("/"))
	//
	//	w := httptest.NewRecorder()
	//	req, _ := http.NewRequest("POST", "/webhook", bytes.NewBuffer([]byte(`{"action": "closed", "github": {"number": 1, "title": "Test github", "body": "Test body"}}`)))
	//	router.ServeHTTP(w, req)
	//
	//	assert.Equal(t, http.StatusOK, w.Code)
	//})
}
