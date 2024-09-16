package Controller

import (
	"chat/Common"
	"chat/Handler/User/Commands/Create"
	"chat/Handler/User/Commands/Login"
	"chat/Redis"
	"chat/Repositories"
	"chat/Repositories/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

// GetUser Login godoc
// @Summary User Login
// @Description Logs in a user with account and password credentials
// @Tags Login
// @Accept  json
// @Produce  json
// @Param LoginRequest body Login.LoginRequest true "Login credentials"
// @Success 200 {object} string "Successfully jwt"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Router /user/Login [post]
func (ctrl UserController) GetUser(c *gin.Context) {
	database := Repositories.GormUserRepository{}.InitDatabase()

	// 初始化 UserRepository
	userRepo := models.NewGormUserRepository(database)

	// 初始化 RedisClient
	redis := Redis.NewRedisService()

	// 初始化 Crypto
	helper := &Common.CryptoHelper{}

	// 初始化 Jwt
	jwt := Common.Jwt{}

	// 注入到 LoginHandler
	loginHandler := Login.NewLoginHandler(userRepo, redis, helper, jwt)

	// 呼叫 業務邏輯
	loginHandler.LoginQueryHandler(c)
}

// CreateUser Login godoc
// @Summary Create User
// @Tags Login
// @Accept  json
// @Produce  json
// @Param UserCreateRequest body Create.UserCreateRequest true "UserCreate Data"
// @Success 200 {object} string "Successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Created User Failed"
// @Router /user/Create [post]
func (ctrl UserController) CreateUser(c *gin.Context) {
	database := Repositories.GormUserRepository{}.InitDatabase()

	// 初始化 UserRepository
	userRepo := models.NewGormUserRepository(database)

	// 初始化 Crypto
	helper := &Common.CryptoHelper{}

	// 注入到 NewLoginHandler
	loginHandler := Create.NewLoginHandler(userRepo, helper)

	// 呼叫 業務邏輯
	loginHandler.CreatUserCommand(c)
}
