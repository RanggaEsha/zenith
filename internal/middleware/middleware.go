package middleware

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Middleware provides the necessary dependencies for handling requests, including database, Redis client, and account repository.
type Middleware struct {
	db    *gorm.DB
	redis *redis.Client
}

// New initializes and returns a new Middleware struct with the provided database, Redis client, and account repository.
func New(db *gorm.DB, redis *redis.Client) *Middleware {
	return &Middleware{
		db:    db,
		redis: redis,
	}
}

//
//// Middleware is a Gin middleware for validating access tokens from Authorization headers and setting authorized account context.
//func Middleware(db *gorm.DB, redisClient *redis.Client) gin.HandlerFunc {
//	repo := repository.NewAccountRepository(db, redisClient)
//	return func(ctx *gin.Context) {
//		resp := common.Response{}
//		if account, err := validateAndExtractAccount(ctx, repo); err != nil {
//			resp.Unauthorized(ctx, []utils.IError{}, err.Error())
//			ctx.Abort()
//		} else {
//			ctx.Set("account", account)
//			ctx.Next()
//		}
//	}
//}
//
//// validateAndExtractAccount validates the token and extracts the account information.
//func validateAndExtractAccount(ctx *gin.Context, repo *repository.AccountRepository) (*model.Account, error) {
//	authHeader := ctx.GetHeader("Authorization")
//	if authHeader == "" {
//		return nil, errormessage.ErrMissingAuthorizationHeader
//	}
//
//	tokenString, err := extractToken(authHeader)
//	if err != nil {
//		return nil, err
//	}
//
//	tokenPayload, err := crp.VerifyToken(tokenString, config.PublicKey)
//	if err != nil {
//		return nil, err
//	}
//
//	if tokenPayload.TokenType != "access_token" {
//		return nil, errormessage.ErrInvalidTokenType
//	}
//
//	isTokenBlacklisted, err := repo.IsTokenBlacklisted(tokenPayload.Jti.String())
//	if err != nil {
//		return nil, err
//	} else if isTokenBlacklisted {
//		return nil, errormessage.ErrInvalidAccessToken
//	}
//
//	account, err := repo.Find(tokenPayload.AccountId)
//	if err != nil {
//		return nil, errormessage.ErrCannotFindAuthorizedAccount
//	}
//
//	return account, nil
//}
//
//// extractToken splits the authorization header to extract the token.
//func extractToken(authHeader string) (string, error) {
//	tokenParts := strings.Split(authHeader, " ")
//	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
//		return "", errormessage.ErrInvalidAccessToken
//	}
//
//	return tokenParts[1], nil
//}
