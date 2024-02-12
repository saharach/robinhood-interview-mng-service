package repository

import (
	"errors"
	"main/internal/api/models"
	"main/internal/db/postgres"
	"time"

	"github.com/jinzhu/gorm"
)

func GetInterviewAll(pagination *models.Pagination) ([]*models.InterviewBody, error) {

	var interviews []*models.InterviewBody

	// Count the total number of interviews that meet the specified conditions
	if err := postgres.DB.Table("interviews").
		Where("interviews.record_status = ?", "A").
		Count(&pagination.Total).Error; err != nil {
		return nil, err
	}

	// Select specific columns from the interviews table and the username column from the users table
	err := postgres.DB.
		Table("interviews").
		Select("interviews.*,  user_infos.first_name || ' ' || user_infos.last_name AS create_user_fullname, user_infos.email AS create_user_email").
		Joins("JOIN user_infos ON interviews.create_user = user_infos.id").
		Where("interviews.record_status = ?", "A").
		Offset((pagination.CurrentPage - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&interviews).Error

	if err != nil {
		return nil, err
	}
	return interviews, nil
}

func GetInterviewByID(interview_id int) (*models.InterviewBody, error) {
	var interview *models.InterviewBody
	err := postgres.DB.
		Table("interviews").
		Select("interviews.*,  user_infos.first_name || ' ' || user_infos.last_name AS create_user_fullname, user_infos.email AS create_user_email").
		Joins("JOIN user_infos ON interviews.create_user = user_infos.id").
		Where("interviews.id = ? AND interviews.record_status = ?", interview_id, "A").
		First(&interview).Error

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("Not found")
		}
		return nil, err
	}

	return interview, nil
}

func CreateInterview(interview *models.Interview, user_id int) error {
	interview.Status = "A"
	interview.RecordStatus = "A"
	interview.CreateUser = user_id
	interview.CreateDate = time.Now()
	if err := postgres.DB.Create(interview).Error; err != nil {
		return err
	}
	if err := CreateInterviewLog(interview); err != nil {
		return err
	}
	return nil
}

func UpdateInterview(interview_id int, interview *models.Interview, user_id int) error {
	_interview, err := GetInterviewByID(interview_id)
	if err != nil {
		return err
	}
	_interview.Name = interview.Name
	_interview.Description = interview.Description
	_interview.Status = interview.Status
	_interview.UpdateUser = user_id
	_interview.UpdateDate = time.Now()
	if err := postgres.DB.Save(_interview.Interview).Error; err != nil {
		return err
	}
	if err := CreateInterviewLog(&_interview.Interview); err != nil {
		return err
	}
	*interview = _interview.Interview
	return nil
}

func UpdateInterviewRecordStatus(interview_id int, status string, user_id int) error {
	_interview, err := GetInterviewByID(interview_id)
	if err != nil {
		return err
	}
	_interview.RecordStatus = status
	_interview.UpdateUser = user_id
	_interview.UpdateDate = time.Now()
	if err := postgres.DB.Save(_interview.Interview).Error; err != nil {
		return err
	}
	if err := CreateInterviewLog(&_interview.Interview); err != nil {
		return err
	}
	return nil
}

func GetInterviewCommentAll(interview_id int, pagination *models.Pagination) ([]*models.InterviewCommentBody, error) {

	var interviewscomment []*models.InterviewCommentBody

	// Count the total number of interviews that meet the specified conditions
	if err := postgres.DB.Table("interview_comments").
		Where("interview_comments.interview_id = ? AND interview_comments.record_status = ?", interview_id, "A").
		Count(&pagination.Total).Error; err != nil {
		return nil, err
	}

	// Select specific columns from the interviews table and the username column from the users table
	err := postgres.DB.
		Table("interview_comments").
		Select("interview_comments.*,  user_infos.first_name || ' ' || user_infos.last_name AS create_user_fullname, user_infos.email AS create_user_email").
		Joins("JOIN user_infos ON interview_comments.create_user = user_infos.id").
		Where("interview_comments.interview_id = ? AND interview_comments.record_status = ?", interview_id, "A").
		Offset((pagination.CurrentPage - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&interviewscomment).Error

	if err != nil {
		return nil, err
	}
	return interviewscomment, nil
}

func GetInterviewCommentByID(interview_id int, comment_id int) (*models.InterviewCommentBody, error) {
	var interviewscomment *models.InterviewCommentBody
	err := postgres.DB.
		Table("interview_comments").
		Select("interview_comments.*,  user_infos.first_name || ' ' || user_infos.last_name AS create_user_fullname, user_infos.email AS create_user_email").
		Joins("JOIN user_infos ON interview_comments.create_user = user_infos.id").
		Where("interview_comments.interview_id =? AND interview_comments.id = ? AND interview_comments.record_status = ?", interview_id, comment_id, "A").
		First(&interviewscomment).Error

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, errors.New("Not found")
		}
		return nil, err
	}
	return interviewscomment, nil
}

func CreateInterviewComment(interview_id int, comment *models.InterviewComment, user_id int) error {
	_interview, err := GetInterviewByID(interview_id)
	if err != nil {
		return err
	}
	comment.InterviewID = _interview.ID
	comment.RecordStatus = "A"
	comment.CreateUser = user_id
	comment.CreateDate = time.Now()
	if err := postgres.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func UpdateInterviewComment(interview_id int, comment_id int, comment *models.InterviewComment, user_id int) error {
	_comment, err := GetInterviewCommentByID(interview_id, comment_id)
	if err != nil {
		return err
	}
	if _comment.CreateUser != user_id {
		return errors.New("Forbidden")
	}
	_comment.Text = comment.Text
	_comment.UpdateUser = user_id
	_comment.UpdateDate = time.Now()
	if err := postgres.DB.Save(_comment.InterviewComment).Error; err != nil {
		return err
	}
	*comment = _comment.InterviewComment
	return nil
}

// DeleteInterviewComment deletes a comment by its ID.
func DeleteInterviewComment(interview_id int, comment_id int, user_id int) error {
	_comment, err := GetInterviewCommentByID(interview_id, comment_id)
	if err != nil {
		return err
	}
	if _comment.CreateUser != user_id {
		return errors.New("Forbidden")
	}
	_comment.RecordStatus = "I"
	_comment.UpdateUser = user_id
	_comment.UpdateDate = time.Now()
	if err := postgres.DB.Save(_comment.InterviewComment).Error; err != nil {
		return err
	}
	return nil
}

// id (int): interview_id
func GetInterviewLogAll(interview_id int, pagination *models.Pagination) ([]*models.InterviewLog, error) {
	var interviewlogs []*models.InterviewLog

	// Count the total number of interviews that meet the specified conditions
	if err := postgres.DB.Table("interview_logs").
		Where("interview_logs.interview_id = ?", interview_id).
		Count(&pagination.Total).Error; err != nil {
		return nil, err
	}

	// Select specific columns from the interviews table and the username column from the users table
	err := postgres.DB.
		Table("interview_logs").
		Select("interview_logs.*,  user_infos.first_name || ' ' || user_infos.last_name AS create_user_fullname, user_infos.email AS create_user_email").
		Joins("JOIN user_infos ON interview_logs.create_user = user_infos.id").
		Where("interview_logs.interview_id = ?", interview_id).
		Order("interview_logs.log_date desc").
		Offset((pagination.CurrentPage - 1) * pagination.PerPage).
		Limit(pagination.PerPage).
		Find(&interviewlogs).Error

	if err != nil {
		return nil, err
	}

	return interviewlogs, nil
}

func CreateInterviewLog(interview *models.Interview) error {

	_log := &models.InterviewLog{
		InterviewID:   interview.ID,
		Name:          interview.Name,
		Description:   interview.Description,
		Status:        interview.Status,
		DefaultSchema: interview.DefaultSchema,
		LogDate:       time.Now(),
	}
	if err := postgres.DB.Create(_log).Error; err != nil {
		return err
	}
	return nil
}
