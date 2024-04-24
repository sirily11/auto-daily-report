package middlewares

import (
	"auto-daily-report/src/config/constants/environments"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// HTTPException represents an HTTP error with a status code and a detail message.
type HTTPException struct {
	StatusCode int
	Detail     string
}

// Error implements the error interface for HTTPException.
func (e *HTTPException) Error() string {
	return fmt.Sprintf("status %d: %s", e.StatusCode, e.Detail)
}

func verifySignature(payloadBody, secretToken, signatureHeader string) error {
	if signatureHeader == "" {
		return &HTTPException{StatusCode: http.StatusForbidden, Detail: "x-hub-signature-256 header is missing!"}
	}

	hash := hmac.New(sha256.New, []byte(secretToken))
	_, err := hash.Write([]byte(payloadBody))
	if err != nil {
		return err
	}
	expectedSignature := "sha256=" + hex.EncodeToString(hash.Sum(nil))

	if !hmac.Equal([]byte(expectedSignature), []byte(signatureHeader)) {
		return &HTTPException{StatusCode: http.StatusForbidden, Detail: "Request signatures didn't match!"}
	}

	return nil
}

func GitHubMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.GetHeader("X-Hub-Signature-256")
		secret := environments.AdminApiKey
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
		payload := buf.String()

		err = verifySignature(payload, secret, signature)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.Next()
	}
}
