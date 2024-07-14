package util

import (
	"fmt"
	"golangliveprojects/iplplayers/pkg/v1/responses"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// DownloadCSV sends StatusOK, formatted message and data.
//
// Example 1: ginutil.JSON(c, &userInfo, "login success")
// Example 2: ginutil.JSON(c, &userDetails, "Email has been sent to: %s", email)
// Example 3: ginutil.JSON(c, &details, "Email sent to: %s, token will expire in %d minutes", email, expires)
func DownloadCSV(c *gin.Context, data interface{}, format string, v ...interface{}) {
	dataContent := fmt.Sprintf("%v", data)
	contentLength := int64(len(dataContent))
	contentType := "application/text/csv"
	contentIoReader := strings.NewReader(dataContent)
	extraHeaders := map[string]string{
		"Content-Disposition": `attachment; filename="output.csv"`,
	}
	c.DataFromReader(http.StatusOK, contentLength, contentType, contentIoReader, extraHeaders)
}

// JSON sends StatusOK, formatted message and data.
//
// Example 1: ginutil.JSON(c, &userInfo, "login success")
// Example 2: ginutil.JSON(c, &userDetails, "Email has been sent to: %s", email)
// Example 3: ginutil.JSON(c, &details, "Email sent to: %s, token will expire in %d minutes", email, expires)
func JSON(c *gin.Context, resp responses.Response, format string, v ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":    http.StatusOK,
		"message":   fmt.Sprintf(format, v...),
		"error":     false,
		"data":      resp.Data,
		"recordSet": resp.RecordSet,
	})
}

// JSONError sends passed in http status code, formatted message and data.
//
// Example 1: ginutil.JSONError(c, http.StatusNotFound, &userInfo, "login failed, no such user")
func JSONError(c *gin.Context, status int, data interface{}, format string, v ...interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  status,
		"message": fmt.Sprintf(format, v...),
		"error":   true,
		"data":    data,
	})
	c.Abort()
}

// JSONUnauthorized ...
func JSONUnauthorized(c *gin.Context, data interface{}, format string, v ...interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  http.StatusUnauthorized,
		"message": string(format),
		"error":   true,
	})
	c.Abort()
}

// JSONInvalidPath ...
func JSONInvalidPath(c *gin.Context, data interface{}, format string, v ...interface{}) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  http.StatusBadRequest,
		"message": string(format),
		"error":   true,
	})
	c.Abort()
}
