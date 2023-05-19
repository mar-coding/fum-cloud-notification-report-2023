package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/services"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		id, err := services.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		// Set the 'id' in the Gin context for future use
		context.Set("id", id)
		context.Next()
	}
}

func getId(ctx *gin.Context) int {
	id, exists := ctx.Get("id")
	if !exists {
		// Handle the case when 'id' is not found in the context
		ctx.JSON(500, gin.H{"error": "id not found in context"})
		return 0
	}

	// Convert the 'id' to string type
	strId := fmt.Sprintf("%v", id)

	intID, err := strconv.Atoi(strId)
	if err != nil {
		// Handle the error when the 'id' cannot be converted to an integer
		ctx.JSON(500, gin.H{"error": "failed to convert id to integer"})
		return 0
	}
	return intID
}

func HandleAllMailRequests(context *gin.Context, sqlDB *sql.DB) {
	//
	// for route: /reports/mail
	//

	tempId := getId(context)
	results := services.GetMailRequests(tempId, sqlDB)
	if len(results) == 0 {
		context.JSON(404, gin.H{"error": "not found"})
	} else {
		context.JSON(http.StatusOK, results)
	}
}

func HandleMailRequestByID(context *gin.Context, sqlDB *sql.DB) {
	//
	// for route: /reports/mail/{requestId}
	//

	requestID := context.Param("requestId")
	tempId := getId(context)
	results := services.GetMailItemsByRequest(requestID, tempId, sqlDB)
	if len(results) == 0 {
		context.JSON(404, gin.H{"error": "not found"})
	} else {
		context.JSON(http.StatusOK, results)
	}

}

func HandleMailRequestByConfigID(context *gin.Context, sqlDB *sql.DB) {
	//
	// for route: /reports/mail/configs/{configId}
	//

	configID := context.Param("configId")
	temp, err := strconv.Atoi(configID)
	if err != nil {
		return
	}
	tempId := getId(context)
	results := services.GetMailItemsByMailConfigId(temp, tempId, sqlDB)
	if len(results) == 0 {
		context.JSON(404, gin.H{"error": "not found"})
	} else {
		context.JSON(http.StatusOK, results)
	}
}
