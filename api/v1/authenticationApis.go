package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginController is
func LoginController(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"username": "l.zenobi",
			"name":     "Luca",
			"surname":  "Zenobi",
			"year":     1990,
		},
	})
}
