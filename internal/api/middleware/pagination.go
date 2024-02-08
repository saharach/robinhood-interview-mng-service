package middleware

import (
	"fmt"
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
		pagination.FirstPageURL = "/api/v1/users?page=1&per_page=" + strconv.Itoa(perPage)
		pagination.LastPageURL = fmt.Sprintf("/api/v1/users?page=%d&per_page=%d", pagination.LastPage, perPage)

		if currentPage < pagination.LastPage {
			pagination.NextPageURL = fmt.Sprintf("/api/v1/users?page=%d&per_page=%d", currentPage+1, perPage)
		} else {
			pagination.NextPageURL = ""
		}

		if currentPage > 1 {
			pagination.PrevPageURL = fmt.Sprintf("/api/v1/users?page=%d&per_page=%d", currentPage-1, perPage)
		} else {
			pagination.PrevPageURL = ""
		}

		// Set the pagination struct in the context for the controller to use
		c.Set("pagination", pagination)

		// Call the next middleware function
		c.Next()

		pagination.LastPage = int(math.Ceil(float64(pagination.Total) / float64(perPage)))
		// fmt.Println("11")
	}
}
