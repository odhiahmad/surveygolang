package main

import (
	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/apiuser/config"
	"github.com/odhiahmad/apiuser/controller"
	"github.com/odhiahmad/apiuser/middleware"
	"github.com/odhiahmad/apiuser/repository"
	"github.com/odhiahmad/apiuser/service"
	"gorm.io/gorm"
)

var (
	db              *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	rumahRepository repository.RumahRepository = repository.NewRumahRepository(db)

	jwtService   service.JWTService   = service.NewJwtService()
	authService  service.AuthService  = service.NewAuthService(userRepository)
	userService  service.UserService  = service.NewUserService(userRepository)
	rumahService service.RumahService = service.NewRumahService(rumahRepository)

	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	rumahController controller.RumahController = controller.NewRumahController(rumahService, jwtService)
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Static("/image", "./fileupload")

	r.Use(CORSMiddleware())

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}
	// middleware.AuthorizeJWT(jwtService)

	userRoutes := r.Group("api/user")
	{
		userRoutes.POST("/create", userController.CreateUser)
		userRoutes.PUT("/update", userController.UpdateUser)
	}

	rumahRoutes := r.Group("api/rumah", middleware.AuthorizeJWT(jwtService))
	{
		rumahRoutes.POST("/create", rumahController.CreateRumah)
		rumahRoutes.PUT("/update", rumahController.UpdateRumah)
		rumahRoutes.POST("/getById", rumahController.FindById)
		rumahRoutes.POST("/getAll", rumahController.FindAll)
		rumahRoutes.POST("/getAllByKota", rumahController.FindAllByKota)
		rumahRoutes.DELETE("/delete", rumahController.Delete)
		rumahRoutes.GET("/statistik", rumahController.Statistik)
		rumahRoutes.POST("/deleteById", rumahController.DeleteByIds)
	}

	r.Run()
}
