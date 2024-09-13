package handler

import (
	"github.com/arifai/go-modular-monolithic/config"
	"github.com/arifai/go-modular-monolithic/internal/account/api/types"
	"github.com/arifai/go-modular-monolithic/internal/account/domain/service"
	"github.com/arifai/go-modular-monolithic/pkg/common"
	"github.com/arifai/go-modular-monolithic/pkg/core"
	"github.com/arifai/go-modular-monolithic/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AuthHandler is a function to handle the auth handler
func AuthHandler(ctx *gin.Context, db *gorm.DB, config *config.Config) {
	resp := new(common.Response)
	accountService := service.NewAccountAuthService(db, config)
	body, err := utils.ValidateBody[types.AccountAuthRequest](ctx)
	if err != nil {
		resp.Error(ctx, err)
		return
	}

	result, err := accountService.Authorize(body)
	if err != nil {
		resp.Error(ctx, err)
		return
	}

	resp.Authorized(ctx, result)
}

// GetAccountHandler is a function to handle the get account handler
func GetAccountHandler(ctx *gin.Context, db *gorm.DB, config *config.Config) {
	resp := new(common.Response)
	accountService := service.NewAccountService(db, config)
	context := core.NewContext(ctx)
	result, err := accountService.GetAccount(context)
	if err != nil {
		resp.Error(ctx, err)
		return
	}

	resp.Success(ctx, result)
}