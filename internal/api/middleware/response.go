package middleware

import (
	"main/internal/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Next()

		if len(c.Errors) > 0 {
			// If there was an error, construct an API response with the real error messages
			var errors []string
			for _, err := range c.Errors {
				errors = append(errors, err.Error())
			}
			apiResponse := &models.APIResponse{
				Success:    false,
				Errors:     errors,
				StatusCode: http.StatusInternalServerError, // Set appropriate status code for client errors
			}
			c.JSON(http.StatusInternalServerError, apiResponse)
		} else if c.Keys["validate"] != nil {
			// If there was an error in the controller, return an API response with the error message
			apiResponse := &models.APIResponse{
				Success:    false,
				Errors:     []string{c.Keys["validate"].(string)},
				StatusCode: http.StatusBadRequest, // Set appropriate status code for internal server errors
			}
			c.JSON(http.StatusBadRequest, apiResponse)
		} else if c.Keys["error"] != nil {
			// If there was an error in the controller, return an API response with the error message
			apiResponse := &models.APIResponse{
				Success:    false,
				Errors:     []string{c.Keys["error"].(string)},
				StatusCode: http.StatusInternalServerError, // Set appropriate status code for internal server errors
			}
			c.JSON(http.StatusInternalServerError, apiResponse)
		} else if c.Keys["data"] != nil {
			// If there were no errors, return an API response with the data and pagination metadata
			apiResponse := &models.APIResponse{
				Success: true,
				Data: map[string]interface{}{
					"meta":  c.Keys["meta"],
					"items": c.Keys["data"],
				},
				StatusCode: http.StatusOK,
			}
			c.JSON(http.StatusOK, apiResponse)
		} else {
			// If there was no data and no errors, return an empty response with a 204 status code
			c.Status(http.StatusNoContent)
		}
	}
}

func NotFound(c *gin.Context) {
	// Return a 404 error in the API response format
	apiResponse := &models.APIResponse{
		Success: false,
		Errors:  []string{"Not found"},
	}
	c.JSON(http.StatusNotFound, apiResponse)
}
