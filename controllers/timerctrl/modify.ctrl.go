package timerctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/subject"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
	"pentag.kr/distimer/utils/notify"
)

func ModifyTimer(c *fiber.Ctx) error {
	data := new(timerMetadataDTO)
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

	subjectID := uuid.MustParse(data.SubjectID)
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	subjectObj, err := dbConn.Subject.Query().Where(subject.ID(subjectID)).Only(context.Background())
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

	timerObj, err := dbConn.Timer.Query().Where(timer.UserID(userID)).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Timer not found",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	var sharedGroupIDs []uuid.UUID
	for _, groupIDStr := range data.SharedGroupIDs {
		groupID, err := uuid.Parse(groupIDStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid group ID",
			})
		}

		exist, err := dbConn.Affiliation.Query().Where(affiliation.And(affiliation.GroupID(groupID), affiliation.UserID(userID))).Exist(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		if !exist {
			return c.Status(403).JSON(fiber.Map{
				"error": "You are not the member of the group",
			})
		}
		sharedGroupIDs = append(sharedGroupIDs, groupID)
	}

	timerObj, err = timerObj.Update().
		Where(timer.UserID(userID)).
		SetContent(data.Content).
		SetSubject(subjectObj).
		ClearSharedGroup().
		AddSharedGroupIDs(sharedGroupIDs...).
		Save(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	go notify.SendTimerUpdate(userID.String(), timerObj, subjectObj)
	return c.JSON(timerDTO{
		ID:             timerObj.ID.String(),
		SubjectID:      data.SubjectID,
		Content:        timerObj.Content,
		SharedGroupIDs: data.SharedGroupIDs,
	})
}
