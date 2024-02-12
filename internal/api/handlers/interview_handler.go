package handlers

import (
	"main/internal/api/models"
	"main/internal/api/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllInterviews(c *gin.Context) {
	pagination, _ := c.MustGet("pagination").(*models.Pagination)
	interview, err := repository.GetInterviewAll(pagination)
	if err != nil {
		c.Set("error", err.Error())
		return
	}
	c.Set("data", interview)
	c.Set("meta", pagination)
}

func GetInterviewByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Parse the string into an integer
	if err != nil {
		c.Set("validate", "Interview ID must be integer.")
		return
	}
	interview, err := repository.GetInterviewByID(id)
	if err != nil {
		if err.Error() == "Not found" {
			c.Set("notfound", "Not found interview.")
			return
		}
		c.Set("error", err.Error())
		return
	}
	c.Set("data", interview)
}
func CreateInterview(c *gin.Context) {
	userInfo, _ := c.Get("user")
	user, _ := userInfo.(*models.Claims)

	var interview models.Interview
	if err := c.ShouldBindJSON(&interview); err != nil {
		c.Set("validate", err.Error())
		return
	}
	if err := repository.CreateInterview(&interview, user.UserId); err != nil {
		c.Set("error", err.Error())
		return
	}
	c.Set("data", interview)

}
func UpdateInterview(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Set("validate", "Interview ID must be integer")
		return
	}

	userInfo, _ := c.Get("user")
	user, _ := userInfo.(*models.Claims)

	var interview *models.Interview
	if err := c.ShouldBindJSON(&interview); err != nil {
		c.Set("validate", err.Error())
		return
	}

	updateInterview := &models.Interview{
		Name:        interview.Name,
		Description: interview.Description,
		Status:      interview.Status,
	}
	err = repository.UpdateInterview(id, updateInterview, user.UserId)
	if err != nil {
		if err.Error() == "Not found" {
			c.Set("notfound", "Not found interview.")
			return
		}
		c.Set("error", err.Error())
		return
	}
	c.Set("data", updateInterview)
}
func DeleteInterview(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Parse the string into an integer
	if err != nil {
		c.Set("validate", "Interview ID must be integer")
		return
	}
	userInfo, _ := c.Get("user")
	user, _ := userInfo.(*models.Claims)

	err = repository.UpdateInterviewRecordStatus(id, "I", user.UserId)
	if err != nil {
		if err.Error() == "Not found" {
			c.Set("notfound", "Not found interview.")
			return
		}
		c.Set("error", err.Error())
		return
	}
	c.Set("data", gin.H{})

}

func ArchiveInterview(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr) // Parse the string into an integer
	if err != nil {
		c.Set("validate", "Interview ID must be integer")
		return
	}
	userInfo, _ := c.Get("user")
	user, _ := userInfo.(*models.Claims)

	err = repository.UpdateInterviewRecordStatus(id, "P", user.UserId)
	if err != nil {
		if err.Error() == "Not found" {
			c.Set("notfound", "Not found interview.")
			return
		}
		c.Set("error", err.Error())
		return
	}
	c.Set("data", gin.H{})
}

func GetInterviewCommentAll(c *gin.Context) {
	pagination, _ := c.MustGet("pagination").(*models.Pagination)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Set("validate", "The interview ID must be an integer.")
		return
	}
	InterviewComment, err := repository.GetInterviewCommentAll(id, pagination)
	if err != nil {
		c.Set("error", err.Error())
		return
	}

	c.Set("data", InterviewComment)
	c.Set("meta", pagination)
}

func CreateInterviewComment(c *gin.Context) {
	userInfo, _ := c.Get("user")
	user, _ := userInfo.(*models.Claims)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Set("validate", "The interview ID must be an integer.")
		return
	}

	var comment models.InterviewComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.Set("validate", err.Error())
		return
	}
	if err := repository.CreateInterviewComment(id, &comment, user.UserId); err != nil {
		if err.Error() == "Not found" {
			c.Set("notfound", "Not found interview.")
			return
		}
		c.Set("error", err.Error())
		return
	}
	c.Set("data", comment)
}

func UpdateInterviewComment(c *gin.Context) {
	userInfo, _ := c.Get("user")
	user, _ := userInfo.(*models.Claims)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Set("validate", "The interview comment ID must be an integer.")
		return
	}
	cidStr := c.Param("cid")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		c.Set("validate", "The interview comment ID must be an integer.")
		return
	}

	var InterviewComment *models.InterviewComment
	if err := c.ShouldBindJSON(&InterviewComment); err != nil {
		c.Set("validate", err.Error())
		return
	}

	updateInterviewComment := &models.InterviewComment{
		Text: InterviewComment.Text,
	}
	err = repository.UpdateInterviewComment(id, cid, updateInterviewComment, user.UserId)
	if err != nil {
		if err.Error() == "Forbidden" {
			c.Set("forbidden", "This user can not edit the comment.")
			return
		}
		if err.Error() == "Not found" {
			c.Set("notfound", "Not found comment.")
			return
		}
		c.Set("error", err.Error())
		return
	}
	c.Set("data", updateInterviewComment)
}

func DeleteInterviewComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Set("validate", "The interview comment ID must be an integer.")
		return
	}
	cidStr := c.Param("cid")
	cid, err := strconv.Atoi(cidStr) // Parse the string into an integer
	if err != nil {
		c.Set("validate", "Interview ID must be integer")
		return
	}
	userInfo, _ := c.Get("user")
	user, _ := userInfo.(*models.Claims)

	err = repository.DeleteInterviewComment(id, cid, user.UserId)
	if err != nil {
		if err.Error() == "Forbidden" {
			c.Set("forbidden", "This user can not edit the comment.")
			return
		}
		c.Set("error", err.Error())
		return
	}
	c.Set("data", gin.H{})
}

func GetInterviewLogAll(c *gin.Context) {
	pagination, _ := c.MustGet("pagination").(*models.Pagination)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Set("validate", "The interview ID must be an integer.")
		return
	}
	InterviewComment, err := repository.GetInterviewLogAll(id, pagination)
	if err != nil {
		c.Set("error", err.Error())
		return
	}

	c.Set("data", InterviewComment)
	c.Set("meta", pagination)
}
