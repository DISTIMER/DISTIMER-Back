package groupctrl

import (
	"context"
	"time"
	"unicode/utf8"

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
	Nickname   string `json:"nickname" example:"nickname between 1 and 20"`
}

// @Summary Join Group with Invite Code
// @Tags Group
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body joinReq true "joinReq"
// @Success 200 {object} groupDTO
// @Router /group/join [post]
func JoinGroup(c *fiber.Ctx) error {
	data := new(joinReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}
	if utf8.RuneCountInString(data.Nickname) < 1 || utf8.RuneCountInString(data.Nickname) > 20 {
		data.Nickname = ""
		return c.Status(400).JSON(fiber.Map{
			"error": "Nickname length should be between 1 and 20",
		})
	}

	dbConn := db.GetDBClient()

	inviteCodeObj, err := dbConn.InviteCode.Query().Where(invitecode.CodeEQ(data.InviteCode)).WithGroup().Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Invite code is not exist",
			})
		}
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	userID := middlewares.GetUserIDFromMiddleware(c)
	affiliationExist, err := dbConn.Affiliation.Query().
		Where(affiliation.And(affiliation.GroupID(inviteCodeObj.Edges.Group.ID), affiliation.UserID(userID))).
		Exist(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if affiliationExist {
		return c.Status(409).JSON(fiber.Map{
			"error": "You are already the member of the group",
		})
	}
	_, err = dbConn.Affiliation.Create().
		SetGroupID(inviteCodeObj.Edges.Group.ID).
		SetUserID(userID).
		SetNickname(data.Nickname).
		SetRole(0).
		Save(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = dbConn.InviteCode.UpdateOne(inviteCodeObj).AddUsed(1).Exec(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	ownerAffiliationObj, err := dbConn.Affiliation.Query().Where(
		affiliation.And(
			affiliation.GroupID(inviteCodeObj.Edges.Group.ID),
			affiliation.Role(2),
		),
	).Only(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	result := groupDTO{
		ID:             inviteCodeObj.Edges.Group.ID.String(),
		Name:           inviteCodeObj.Edges.Group.Name,
		Description:    inviteCodeObj.Edges.Group.Description,
		NicknamePolicy: inviteCodeObj.Edges.Group.NicknamePolicy,
		RevealPolicy:   inviteCodeObj.Edges.Group.RevealPolicy,
		InvitePolicy:   inviteCodeObj.Edges.Group.InvitePolicy,
		CreateAt:       inviteCodeObj.Edges.Group.CreatedAt.Format(time.RFC3339),
		OwnerNickname:  ownerAffiliationObj.Nickname,
	}

	return c.JSON(result)
}
