package middleware

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/config"
	"github.com/mar-coding/fum-cloud-notification-report-2023/app/models"
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

func getBaseURL(u *url.URL) string {
	config := config.LoadFromEnv()
	var base string = config.AppUrl

	baseURL := &url.URL{
		Scheme: u.Scheme,
		Host:   u.Host,
		Path:   u.Path,
	}
	fullURL := base + baseURL.String()
	return fullURL
}

func createPageURL(baseURL string, page int, pageSize int) *url.URL {
	u, _ := url.Parse(baseURL)
	q := u.Query()
	q.Set("page", strconv.Itoa(page))
	q.Set("pageSize", strconv.Itoa(pageSize))
	u.RawQuery = q.Encode()

	return u
}

func checkUrl(totalItems int, pageSize int, pageNum int) bool {
	totalPages := removeDecimal(float64(totalItems) / float64(pageSize))
	if pageNum <= totalPages {
		return true
	}
	return false
}

func createHATEOASLink(sqlDB *sql.DB, context *gin.Context, pageNum int, pageSize int) models.HATEOASLinks {
	baseURL := getBaseURL(context.Request.URL)
	selfURL := createPageURL(baseURL, pageNum, pageSize)
	selfLink := selfURL.String()

	links := models.HATEOASLinks{
		Self: selfLink,
	}

	nextPageNum := pageNum + 1
	nextPageURL := createPageURL(baseURL, nextPageNum, pageSize)
	nextLink := nextPageURL.String()

	totalItems := len(services.GetMailRequests(getId(context), sqlDB))
	if checkUrl(totalItems, pageSize, nextPageNum) {
		links.Next = nextLink
	}

	prevPageNum := pageNum - 1
	prevPageURL := createPageURL(baseURL, prevPageNum, pageSize)
	prevLink := prevPageURL.String()

	if pageNum > 1 {
		links.Prev = prevLink
	}

	return links
}

func removeDecimal(num float64) int {
	if float64(int(num)) == num {
		return int(num)
	} else {
		return int(num) + 1
	}
}

func HATEOASOutputEmail(sqlDB *sql.DB, context *gin.Context, data []models.OutputEmail, pageNum int, pageSize int, totalItems int) models.PaginationResponseOutputEmail {
	links := createHATEOASLink(sqlDB, context, pageNum, pageSize)
	totalPages := removeDecimal(float64(totalItems) / float64(pageSize))

	var response models.PaginationResponseOutputEmail
	response = models.PaginationResponseOutputEmail{
		Meta: models.HATEOASMetaData{
			Page:       pageNum,
			PageSize:   pageSize,
			TotalPages: totalPages,
			TotalItems: totalItems,
			Links:      links,
		},
		Data: data,
	}
	return response
}

func HATEOASOutputReq(sqlDB *sql.DB, context *gin.Context, data []models.OutputReq, pageNum int, pageSize int, totalItems int) models.PaginationResponseOutputReq {
	links := createHATEOASLink(sqlDB, context, pageNum, pageSize)
	totalPages := removeDecimal(float64(totalItems) / float64(pageSize))

	var response models.PaginationResponseOutputReq
	response = models.PaginationResponseOutputReq{
		Meta: models.HATEOASMetaData{
			Page:       pageNum,
			PageSize:   pageSize,
			TotalPages: totalPages,
			TotalItems: totalItems,
			Links:      links,
		},
		Data: data,
	}
	return response
}

func HandleAllMailRequests(context *gin.Context, sqlDB *sql.DB) {
	//
	// for route: /reports/mail
	// or route: /reports/main?page=x&pageSize=y for pagination
	//

	var has_problem bool = false
	pageNumber := context.Query("page")
	pageSize := context.Query("pageSize")
	totalItems := len(services.GetMailRequests(getId(context), sqlDB))

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
			response := HATEOASOutputEmail(sqlDB, context, results, pageNum, pageSz, totalItems)
			context.JSON(http.StatusOK, response)
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
	totalItems := len(services.GetMailRequests(getId(context), sqlDB))

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
			response := HATEOASOutputReq(sqlDB, context, results, pageNum, pageSz, totalItems)
			context.JSON(http.StatusOK, response)
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
	totalItems := len(services.GetMailRequests(getId(context), sqlDB))

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
			response := HATEOASOutputReq(sqlDB, context, results, pageNum, pageSz, totalItems)
			context.JSON(http.StatusOK, response)
		}
	}
}
