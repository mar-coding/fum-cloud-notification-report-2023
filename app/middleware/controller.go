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
		// Set the 'id' in the Gin context
		context.Set("id", id)
		context.Next()
	}
}

func getId(ctx *gin.Context) int {
	id, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(500, gin.H{"error": "id not found in context"})
		return 0
	}

	strId := fmt.Sprintf("%v", id)

	intID, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "failed to convert id to integer"})
		return 0
	}
	return intID
}

func HandleAllMailRequests(context *gin.Context, sqlDB *sql.DB) {
	//
	// for route: /reports/mail
	// or route: /reports/main?page=x&pageSize=y for pagination
	//

	var has_problem bool = false
	pageNumber := context.Query("page")
	pageSize := context.Query("pageSize")

	// Check if page number and page size are provided
	if pageNumber == "" || pageSize == "" {
		// Pagination metadata is not provided, return all items
		tempId := getId(context)
		results := services.GetMailRequests(tempId, sqlDB)
		if len(results) == 0 {
			context.JSON(404, gin.H{"error": "not found"})
		} else {
			context.JSON(http.StatusOK, results)
		}
	} else {
		// Convert page number and page size to integers
		pageNum, err := strconv.Atoi(pageNumber)
		if err != nil {
			context.JSON(400, gin.H{"error": "bad request"})
			has_problem = true
		}

		pageSz, err := strconv.Atoi(pageSize)
		if err != nil {
			if !has_problem {
				context.JSON(400, gin.H{"error": "bad request"})
				has_problem = true
			}
		}
		offset := (pageNum - 1) * pageSz
		tempId := getId(context)

		results := services.GetPaginatedMailRequests(tempId, offset, pageSz, sqlDB)

		if len(results) == 0 {
			if !has_problem {
				context.JSON(404, gin.H{"error": "not found"})
			}
		} else {
			context.JSON(http.StatusOK, results)
		}
	}

}

func HandleMailRequestByID(context *gin.Context, sqlDB *sql.DB) {
	//
	// for route: /reports/mail/{requestId}
	// or route: /reports/mail/{requestId}?page=x&pageSize=y for pagination
	//

	var has_problem bool = false
	pageNumber := context.Query("page")
	pageSize := context.Query("pageSize")

	if pageNumber == "" || pageSize == "" {
		// Pagination metadata is not provided, return all items
		requestID := context.Param("requestId")
		tempId := getId(context)
		results := services.GetMailItemsByRequest(requestID, tempId, sqlDB)
		if len(results) == 0 {
			context.JSON(404, gin.H{"error": "not found"})
		} else {
			context.JSON(http.StatusOK, results)
		}
	} else {
		// Convert page number and page size to integers
		pageNum, err := strconv.Atoi(pageNumber)
		if err != nil {
			context.JSON(400, gin.H{"error": "bad request"})
			has_problem = true
		}

		pageSz, err := strconv.Atoi(pageSize)
		if err != nil {
			if !has_problem {
				context.JSON(400, gin.H{"error": "bad request"})
				has_problem = true
			}
		}
		offset := (pageNum - 1) * pageSz
		requestID := context.Param("requestId")
		tempId := getId(context)

		results := services.GetPaginatedMailItemsByRequest(requestID, tempId, offset, pageSz, sqlDB)

		if len(results) == 0 {
			if !has_problem {
				context.JSON(404, gin.H{"error": "not found"})
			}
		} else {
			context.JSON(http.StatusOK, results)
		}
	}
}

func HandleMailRequestByConfigID(context *gin.Context, sqlDB *sql.DB) {
	//
	// for route: /reports/mail/configs/{configId}
	// or route: /reports/mail/configs/{configId}?page=x&pageSize=y for pagination
	//

	var has_problem bool = false
	pageNumber := context.Query("page")
	pageSize := context.Query("pageSize")

	if pageNumber == "" || pageSize == "" {
		// Pagination metadata is not provided, return all items
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
	} else {
		// Convert page number and page size to integers
		pageNum, err := strconv.Atoi(pageNumber)
		if err != nil {
			context.JSON(400, gin.H{"error": "bad request"})
			has_problem = true
		}

		pageSz, err := strconv.Atoi(pageSize)
		if err != nil {
			if !has_problem {
				context.JSON(400, gin.H{"error": "bad request"})
				has_problem = true
			}
		}
		offset := (pageNum - 1) * pageSz
		configID := context.Param("configId")
		temp, err := strconv.Atoi(configID)
		if err != nil {
			return
		}
		tempId := getId(context)

		results := services.GetPaginatedMailItemsByMailConfigId(temp, tempId, offset, pageSz, sqlDB)

		if len(results) == 0 {
			if !has_problem {
				context.JSON(404, gin.H{"error": "not found"})
			}
		} else {
			context.JSON(http.StatusOK, results)
		}
	}
}
