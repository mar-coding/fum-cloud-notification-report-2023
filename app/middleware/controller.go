package middleware

import (
	"database/sql"
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
		err := services.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}

func HandleAllMailRequests(context *gin.Context, sqlDB *sql.DB) {
	//
	// for route: /reports/mail
	//

	results := services.GetMailRequests(2, sqlDB)
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
	results := services.GetMailItemsByRequest(requestID, 2, sqlDB)
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
	results := services.GetMailItemsByMailConfigId(temp, 2, sqlDB)
	if len(results) == 0 {
		context.JSON(404, gin.H{"error": "not found"})
	} else {
		context.JSON(http.StatusOK, results)
	}
}
