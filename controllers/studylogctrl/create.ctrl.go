package studylogctrl

import (
	"context"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/configs"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/studylog"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type createStudyLogReq struct {
	SubjectID     string   `json:"subject_id" validate:"required,uuid" example:"subject_id"`
	StartAt       string   `json:"start_at" validate:"required" example:"2020-08-28T09:20:26.187+09:00"`
	EndAt         string   `json:"end_at" validate:"required" example:"2020-08-28T09:20:26.187+09:00"`
	Content       string   `json:"content" example:"content between 0 and 30"`
	GroupsToShare []string `json:"groups_to_share" validate:"required" example:"group_id"`
}

func CreateStudyLog(c *fiber.Ctx) error {

	// Data Parsing
	userID := middlewares.GetUserIDFromMiddleware(c)

	data := new(createStudyLogReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	if utf8.RuneCountInString(data.Content) > 30 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Content length should be between 1 and 30",
		})
	}
	// parse date with rf3339 format
	startAt, err := time.Parse(time.RFC3339, data.StartAt)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid start_at",
		})
	}
	endAt, err := time.Parse(time.RFC3339, data.EndAt)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid end_at",
		})
	}
	if startAt.After(endAt) {
		return c.Status(400).JSON(fiber.Map{
			"error": "start_at should be before end_at",
		})
	}
	if endAt.After(time.Now()) {
		return c.Status(400).JSON(fiber.Map{
			"error": "end_at should be before now",
		})
	}
	// if more than 24 hours return error
	if endAt.Sub(startAt) > time.Hour*24 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Study log should be less than 24 hours",
		})
	}

	// truncate milliseconds
	startAt = startAt.Truncate(time.Second)
	endAt = endAt.Truncate(time.Second)

	subjectID, err := uuid.Parse(data.SubjectID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid subject ID",
		})
	}

	var groupIDs []uuid.UUID
	count := 1
	for _, groupIDStr := range data.GroupsToShare {
		if count > configs.FreePlanGroupJoinLimit {
			return c.Status(400).JSON(fiber.Map{
				"error": "You can't share more than limit of free plan groups",
			})
		}

		id, err := uuid.Parse(groupIDStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid group ID",
			})
		}
		groupIDs = append(groupIDs, id)

		count++
	}

	// DB Operation
	dbConn := db.GetDBClient()

	// check if study log is already exist at the same time
	exist, err := dbConn.StudyLog.Query().
		Where(studylog.And(
			studylog.HasUserWith(user.ID(userID)),
			studylog.StartAtLTE(endAt),
			studylog.EndAtGTE(startAt))).
		Exist(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if exist {
		return c.Status(409).JSON(fiber.Map{
			"error": "Study log is already exist at the same time",
		})
	}

	subjectObj, err := dbConn.Subject.Get(context.Background(), subjectID)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Subject not found",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	categoryObj, err := subjectObj.QueryCategory().WithUser().Only(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if categoryObj.Edges.User.ID != userID {
		return c.Status(403).JSON(fiber.Map{
			"error": "You are not the owner of the subject",
		})
	}

	for _, groupID := range groupIDs {
		exist, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Exist(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		} else if !exist {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not the member of the group: " + groupID.String(),
			})
		}
	}

	_, err = dbConn.StudyLog.Create().
		SetSubject(subjectObj).
		SetStartAt(startAt).
		SetEndAt(endAt).
		SetContent(data.Content).
		SetUserID(userID).
		AddSharedGroupIDs(groupIDs...).
		Save(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	return c.JSON(myStudyLogDTO{
		ID:            uuid.New().String(),
		SubjectID:     subjectID.String(),
		StartAt:       startAt.Format(time.RFC3339),
		EndAt:         endAt.Format(time.RFC3339),
		Content:       data.Content,
		GroupsToShare: data.GroupsToShare,
	})

}
