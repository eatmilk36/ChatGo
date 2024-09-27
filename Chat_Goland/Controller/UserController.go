package Controller

import (
	"Chat_Goland/Handler/User/Commands/Create"
	"Chat_Goland/Handler/User/Commands/Login"
	"Chat_Goland/Interface"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepo      Interface.UserRepositoryInterface
	redisService  Interface.RedisServiceInterface
	cryptoService Interface.CryptoServiceInterService
	jwtService    Interface.JwtServiceInterface
	logService    Interface.LogServiceInterface
}

func NewUserController(userRepo Interface.UserRepositoryInterface, redis Interface.RedisServiceInterface, helper Interface.CryptoServiceInterService, jwt Interface.JwtServiceInterface, log Interface.LogServiceInterface) *UserController {
	return &UserController{
		userRepo:      userRepo,
		redisService:  redis,
		cryptoService: helper,
		jwtService:    jwt,
		logService:    log,
	}
}

// GetUser Login godoc
// @Summary Model Login
// @Description Logs in a user with account and password credentials
// @Tags Login
// @Accept  json
// @Produce  json
// @Param LoginRequest body Login.LoginRequest true "Login credentials"
// @Success 200 {object} string "Successfully jwt"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Model not found"
// @Router /user/Login [post]
func (ctrl UserController) GetUser(c *gin.Context) {
	Login.NewLoginHandler(ctrl.userRepo, ctrl.redisService, ctrl.cryptoService, ctrl.jwtService, ctrl.logService).LoginQueryHandler(c)
}

// CreateUser godoc
// @Summary Create Model
// @Tags Login
// @Accept  json
// @Produce  json
// @Param UserCreateRequest body Create.UserCreateRequest true "UserCreate Data"
// @Success 200 {object} string "Successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 404 {object} map[string]interface{} "Created Model Failed"
// @Router /user/Create [post]
func (ctrl UserController) CreateUser(c *gin.Context) {
	Create.NewLoginHandler(ctrl.userRepo, ctrl.cryptoService).CreatUserCommand(c)
}
