package middleware

import (
	"main/internal/api/models"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetPaginationData() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the current page and per page values from the query string
		currentPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))

		pagination := &models.Pagination{}
		// Update the pagination struct with the new values
		pagination.PerPage = perPage
		pagination.CurrentPage = currentPage

		pagination.FirstPage = 1

		// Set the pagination struct in the context for the controller to use
		c.Set("pagination", pagination)

		// Call the next middleware function
		c.Next()

		pagination.LastPage = int(math.Ceil(float64(pagination.Total) / float64(perPage)))
	}
}
