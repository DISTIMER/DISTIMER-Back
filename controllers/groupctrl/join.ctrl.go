package groupctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/affiliation"
	"pentag.kr/distimer/ent/invitecode"
	"pentag.kr/distimer/middlewares"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type joinReq struct {
	InviteCode string `json:"invite_code" validate:"required,len=7"`
}

type joinRes struct {
	Group *ent.Group `json:"group"`
}

// @Summary Join Group with Invite Code
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body joinReq true "joinReq"
// @Success 200 {object} joinRes
// @Router /group/join [post]
func JoinGroup(c *fiber.Ctx) error {
	data := new(joinReq)
	if err := dto.Bind(c, data); err != nil {
		return err
	}

	dbConn := db.GetDBClient()

	inviteCodeObj, err := dbConn.InviteCode.Query().Where(invitecode.CodeEQ(data.InviteCode)).WithGroup().Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Invite code is not exist",
			})
		}
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	userID := middlewares.GetUserIDFromMiddleware(c)
	affiliationExist, err := dbConn.Affiliation.Query().
		Where(affiliation.And(affiliation.GroupID(inviteCodeObj.Edges.Group.ID), affiliation.UserID(userID))).
		Exist(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if affiliationExist {
		return c.Status(409).JSON(fiber.Map{
			"error": "You are already the member of the group",
		})
	}
	userObj, err := dbConn.User.Get(context.Background(), userID)
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	_, err = dbConn.Affiliation.Create().
		SetGroupID(inviteCodeObj.Edges.Group.ID).
		SetUserID(userID).
		SetNickname(userObj.Name).
		SetRole(0).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = dbConn.InviteCode.UpdateOne(inviteCodeObj).AddUsed(1).Exec(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(joinRes{
		Group: inviteCodeObj.Edges.Group,
	})
}
