package controller

import (
	"lecter/goserver/internal/app/gochat/controller/request"
	"lecter/goserver/internal/app/gochat/controller/response"
	"lecter/goserver/internal/app/gochat/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserProfileController struct{}

var userProfileService = service.UserProfileService{}
var authenicateService = service.AuthenticationService{}

/*
 *	ユーザープロフィールを取得する
 */
func (upc UserProfileController) Select(ctx *gin.Context) {
	// ユーザーID取得
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}
	// ユーザーモデル取得
	model, error := userProfileService.SelectUserProfile(userId)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, *model)
}

/*
 * ユーザープロフィールを新規作成・更新を行う
 */
func (upc UserProfileController) Update(ctx *gin.Context) {
	// 更新対象のユーザーID取得
	userId, err := uuid.Parse(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(response.ValidationError("invalid userId").ToResponse())
		return
	}
	// リクエスト送信者のユーザー名取得
	username, exists := ctx.Get("username")
	if !exists {
		ctx.JSON(response.InternalError("failed to get username").ToResponse())
		return
	}
	// リクエスト送信者と対象のIDの一致確認
	if error := authenicateService.IsUserRelated(userId, username.(string)); error != nil {
		ctx.JSON(error.ToResponse())
		return
	}

	// 更新リクエストボディ取得
	var request request.UserProfileUpdateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(response.ValidationError("invalid request").ToResponse())
		return
	}

	// ユーザープロフィールを更新
	model, error := userProfileService.UpdateUserProfile(userId, request.DisplayName, request.Url, request.Description)
	if error != nil {
		ctx.JSON(error.ToResponse())
		return
	}
	ctx.JSON(http.StatusOK, *model)
}
