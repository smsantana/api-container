package main

import (
	"api-container/infrastructure/auth"
	"api-container/infrastructure/persistence"
	"api-container/interfaces"
	"api-container/interfaces/fileupload"
	"api-container/interfaces/middleware"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// conexao com banco de dados
	services, err := persistence.NewRepositories(DbConnect().dbdriver, DbConnect().user, DbConnect().password, DbConnect().port, DbConnect().host, DbConnect().dbname)
	OnErr(err)

	//fechando conexao com db ao final da execucao
	defer services.Close()
	services.Automigrate()

	//Abrindo conexao com redis
	redisService, err := auth.NewRedisDB(RedisConnect().redis_host, RedisConnect().redis_port, RedisConnect().redis_password)
	OnErr(err)

	tk := auth.NewToken()
	fd := fileupload.NewFileUpload()

	users := interfaces.NewUsers(services.User, redisService.Auth, tk)
	foods := interfaces.NewFood(services.Food, services.User, fd, redisService.Auth, tk)

	authenticate := interfaces.NewAuthenticate(services.User, redisService.Auth, tk)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.POST("/usuarios", users.SaveUser)
	r.GET("/usuarios", users.GetUsers)
	r.GET("/usuarios/:usuar_id", users.GetUser)

	r.POST("/comida", middleware.AuthMiddleware(), middleware.MaxSizeAllowed(8192000), foods.SaveFood)
	r.PUT("/comida/:com_id", middleware.AuthMiddleware(), middleware.MaxSizeAllowed(8192000), foods.UpdateFood)
	r.GET("/comida/:com_id", foods.GetFoodAndCreator)
	r.DELETE("/comida/:com_id", middleware.AuthMiddleware(), foods.DeleteFood)
	r.GET("/comida", foods.GetAllFood)

	r.POST("/login", authenticate.Login)
	r.POST("/logout", authenticate.Logout)
	r.POST("/refresh", authenticate.Refresh)

	app_port := os.Getenv("PORT")
	if app_port == "" {
		app_port = "8888"
	}
	log.Fatal(r.Run(":" + app_port))
}

//Save three lines of code when you need to panic.
func OnErr(err error) {
	if err != nil {
		panic(fmt.Sprintf("Error in function: %s; Details: %s", getFunctionName(), err))
	}
}

func getFunctionName() string {
	// Get the caller of OnErr's name.
	if pc, _, _, ok := runtime.Caller(2); ok {
		return runtime.FuncForPC(pc).Name()
	}
	return "function not found"
}

//CHAOS If you need a chaos monkey in your code to flesh out error handling paths, this is for you
func Chaos(err error) error {
	if rand.Intn(1) == 1 {
		return err
	} else {
		return errors.New("The chaos monkey found you!")
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}
