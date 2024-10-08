package timerctrl

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/timer"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/logger"
	"pentag.kr/distimer/utils/notify"
)

func GetMyTimerInfo(c *fiber.Ctx) error {
	userID := middlewares.GetUserIDFromMiddleware(c)

	dbConn := db.GetDBClient()

	foundTimer, err := dbConn.Timer.Query().Where(timer.HasUserWith(user.ID(userID))).WithSharedGroup().First(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"info": "Timer not found",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if time.Since(foundTimer.StartAt) > time.Hour*24 {
		// delete timer if it's older than 24 hours
		err := dbConn.Timer.DeleteOne(foundTimer).Exec(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
		go notify.SendTimerDelete(userID.String())
		return c.Status(404).JSON(fiber.Map{
			"info": "Timer not found",
		})
	}
	return c.JSON(timerDTO{
		ID:        foundTimer.ID.String(),
		SubjectID: foundTimer.SubjectID.String(),
		Content:   foundTimer.Content,
		StartAt:   foundTimer.StartAt.Format(time.RFC3339),
		SharedGroupIDs: func() []string {
			sharedGroups := foundTimer.Edges.SharedGroup
			result := make([]string, len(sharedGroups))
			for i, sharedGroup := range sharedGroups {
				result[i] = sharedGroup.ID.String()
			}
			return result
		}(),
	})
}
